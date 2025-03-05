package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"resourceup/condition"
	"resourceup/item_model"
	"resourceup/model"
	"resourceup/range_dispatch_type"
	"resourceup/select_type"
	"resourceup/special"
	"strings"
)

type LegacyModel struct {
	Parent    string          `json:"parent,omitempty"`
	Overrides []ModelOverride `json:"overrides,omitempty"`
}

type ModelOverride struct {
	Predicate map[string]float64 `json:"predicate"`
	Model     string             `json:"model"`
}

type ModelParam struct {
	Parent        string
	ModelOverride string

	ModelBase string
}

var rootDir string

func main() {
	flag.StringVar(&rootDir, "dir", ".", "Folder path")
	// 解析命令列參數，指定要處理的資料夾
	flag.Parse()

	// 遍歷資料夾中的所有 JSON 檔案
	abs, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatalf("Error occured：%v", err)
	}

	{
		rPath := filepath.Join(abs, "assets/minecraft/models/**.json")
		matches, err := filepath.Glob(rPath)
		if err != nil {
			log.Fatalf("Error occured：%v", err)
		}

		for _, path := range matches {
			var originalMap map[string]interface{}
			file, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("Error occured：%v", err)
			}
			err = json.Unmarshal(file, &originalMap)
			if err != nil {
				log.Printf("Error occured：%v", err)
			}
		}
	}

	rPath := filepath.Join(abs, "assets/minecraft/models/item/*.json")

	matches, err := filepath.Glob(rPath)
	if err != nil {
		log.Fatalf("Error occured：%v", err)
	}

	for _, path := range matches {
		itemName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		log.Printf("Processing %s", itemName)

		// 讀取 JSON 檔案
		var data LegacyModel
		var originalMap map[string]interface{}
		{
			file, err := os.ReadFile(path)
			if err != nil {
				log.Fatalf("Error occured：%v", err)
			}
			err = json.Unmarshal(file, &data)
			if err != nil {
				log.Printf("Error occured：%v", err)
			}
			err = json.Unmarshal(file, &originalMap)
			if err != nil {
				log.Printf("Error occured：%v", err)
			}
		}

		if len(data.Overrides) == 0 {
			continue
		}

		defaultModel := getModelByPath(ModelParam{
			Parent:        data.Parent,
			ModelBase:     itemName,
			ModelOverride: "item/" + itemName,
		})
		_ = defaultModel

		var items []Item
		items = append(items, Item{
			Predicate: make(map[string]float64),
			Model:     defaultModel,
		})
		for _, override := range data.Overrides {
			normalModel := getModelByPath(ModelParam{
				Parent:        data.Parent,
				ModelOverride: override.Model,
				ModelBase:     itemName,
			})
			item := Item{
				Predicate: override.Predicate,
				Model:     normalModel,
			}
			preprocessItems(&item)
			items = append(items, item)
		}

		tree := buildTree(items, make(map[string]bool))

		model := createModel(tree, 0)
		output := &OutputModel{
			HandAnimationOnSwap: false,
			Model:               model,
		}

		result, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			log.Printf("Error occured：%v", err)
		}

		delete(originalMap, "overrides")

		newOriginalFile, err := json.MarshalIndent(originalMap, "", "  ")
		if err != nil {
			log.Printf("Error occured：%v", err)
		}

		err = os.WriteFile(path, newOriginalFile, os.ModePerm)
		if err != nil {
			log.Printf("Error occured：%v", err)
		}
		dir, file := filepath.Split(path)

		dir = filepath.Clean(filepath.Join(dir, "..", "..", "items"))
		_ = os.MkdirAll(dir, os.ModeDir)
		newPath := filepath.Join(dir, file)

		err = os.WriteFile(newPath, result, os.ModePerm)
		if err != nil {
			log.Printf("Error occured：%v", err)
		}
	}
}

type OutputModel struct {
	HandAnimationOnSwap bool             `json:"hand_animation_on_swap,omitempty"`
	Model               item_model.Model `json:"model"`
}

func createModel(node *Node, layer int) item_model.Model {
	if len(node.SubNodes) == 0 {
		return node.Model
	}
	if len(node.SubNodes) == 1 {
		for key, nodes := range node.SubNodes {
			switch key {
			case "custom_model_data":
				r := &range_dispatch_type.CustomModelData{}
				for _, n := range nodes {
					data := createModel(n, layer+1)
					if n.Predicate.Value == 0 {
						r.Value.Fallback = data
					} else {
						r.Value.Entries = append(r.Value.Entries,
							range_dispatch_type.Value[range_dispatch_type.CustomModelDataValue]{
								Threshold: range_dispatch_type.CustomModelDataValue(n.Predicate.Value),
								Model:     data,
							},
						)
					}
				}
				return &range_dispatch_type.RangeDispatcher{RangeDispatchType: r}
			case "pull":
				r := &range_dispatch_type.CrossbowPull{}
				for _, n := range nodes {
					data := createModel(n, layer+1)
					if n.Predicate.Value == 0 {
						r.Value.Fallback = data
					} else {
						r.Value.Entries = append(r.Value.Entries,
							range_dispatch_type.Value[float64]{
								Threshold: n.Predicate.Value,
								Model:     data,
							},
						)
					}
				}

				return &range_dispatch_type.RangeDispatcher{RangeDispatchType: r}
			case "charge_type":
				s := select_type.ChargeType{}

				for _, n := range nodes {
					data := createModel(n, layer+1)

					switch n.Predicate.Value {
					case 0:
						s.Value.Fallback = data
					case 1:
						s.Value.Case = append(s.Value.Case, select_type.Value[select_type.ChargeTypeValue]{
							When:  []select_type.ChargeTypeValue{select_type.ChargeTypeRocket},
							Model: data,
						})
					case 2:
						s.Value.Case = append(s.Value.Case, select_type.Value[select_type.ChargeTypeValue]{
							When:  []select_type.ChargeTypeValue{select_type.ChargeTypeArrow},
							Model: data,
						})
					}
				}

				return &select_type.Select{SelectType: s}
			case "cast":
				c := &condition.FishingRodCast{}
				for _, n := range nodes {
					if n.Predicate.Value == 0 {
						c.Value.FalseModel = createModel(n, layer+1)
					} else {
						c.Value.TrueModel = createModel(n, layer+1)
					}
				}

				return &condition.Condition{ConditionType: c}
			case "pulling":
				c := &condition.UsingItem{}
				for _, n := range nodes {
					if n.Predicate.Value == 0 {
						c.Value.FalseModel = createModel(n, layer+1)
					} else {
						c.Value.TrueModel = createModel(n, layer+1)
					}
				}

				return &condition.Condition{ConditionType: c}
			case "damage":
				r := &range_dispatch_type.Damage{}

				for _, n := range nodes {
					data := createModel(n, layer+1)
					if n.Predicate.Value == 0 {
						r.Value.Fallback = data
					} else {
						r.Value.Entries = append(r.Value.Entries,
							range_dispatch_type.Value[float64]{
								Threshold: n.Predicate.Value,
								Model:     data,
							},
						)
					}
				}

				return &range_dispatch_type.RangeDispatcher{RangeDispatchType: r}
			case "damaged":
				c := &condition.Damaged{}
				for _, n := range nodes {
					if n.Predicate.Value == 0 {
						c.Value.FalseModel = createModel(n, layer+1)
					} else {
						c.Value.TrueModel = createModel(n, layer+1)
					}
				}

				return &condition.Condition{ConditionType: c}
			case "broken":
				c := &condition.Broken{}
				for _, n := range nodes {
					if n.Predicate.Value == 0 {
						c.Value.FalseModel = createModel(n, layer+1)
					} else {
						c.Value.TrueModel = createModel(n, layer+1)
					}
				}

				return &condition.Condition{ConditionType: c}
			case "cooldown":
				r := &range_dispatch_type.Cooldown{}

				for _, n := range nodes {
					data := createModel(n, layer+1)
					if n.Predicate.Value == 0 {
						r.Value.Fallback = data
					} else {
						r.Value.Entries = append(r.Value.Entries,
							range_dispatch_type.Value[range_dispatch_type.ZeroToOneDataValue]{
								Threshold: range_dispatch_type.ZeroToOneDataValue(n.Predicate.Value),
								Model:     data,
							},
						)
					}
				}

				return &range_dispatch_type.RangeDispatcher{RangeDispatchType: r}
			default:
				fmt.Println("Unknown key", key)
			}
		}

		return nil
	}

	// 有多個分支的情況，使用 Composite
	composite := &model.Composite{}
	//
	for key, nodes := range node.SubNodes {
		_ = nodes
		fmt.Println(key)

		//switch key {
		//case "broken":
		//	for _, n := range nodes {
		//
		//	}
		//}
	}

	return composite
}

// printTree 輔助函數，印出樹狀結構
func printTree(node *Node, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	if node.Predicate != nil {
		fmt.Printf("%s- [%s: %v]\n", indent, node.Predicate.Key, node.Predicate.Value)
	} else {
		fmt.Printf("%s* Root\n", indent)
	}
	if node.Model != nil {
		fmt.Printf("%s  Model: %s\n", indent, node.Model)
	}
	for key, childGroup := range node.SubNodes {
		fmt.Printf("%s  Group: %s\n", indent, key)
		for _, child := range childGroup {
			printTree(child, level+1)
		}
	}
}

// Item 定義了輸入的資料結構
type Item struct {
	Predicate map[string]float64 `json:"predicate"`
	Model     item_model.Model   `json:"model"`
}

type PredicatePair struct {
	Key   string
	Value float64
}

type Node struct {
	Predicate *PredicatePair
	Model     item_model.Model
	// 將 SubNodes 由 slice 轉成以 predicate key 分組的 map，
	// 每個 key 對應一個 []*Node 切片
	SubNodes map[string][]*Node
}

// 優先鍵列表：這些鍵會優先被選用作為分割依據
var priorityKeys = []string{"custom_model_data"}

func preprocessItems(item *Item) {
	// firework, charged merge to charge_type none/rocket/arrow
	if item.Predicate["charged"] == 1 {
		delete(item.Predicate, "charged")
		item.Predicate["charge_type"] = 2
	}

	if item.Predicate["firework"] == 1 {
		delete(item.Predicate, "firework")
		item.Predicate["charge_type"] = 1
	}

}

// buildTree 透過遞迴將 items 分組建立樹，usedKeys 用來記錄已使用過的 predicate 鍵 ，謝謝ChatGPT
func buildTree(items []Item, usedKeys map[string]bool) *Node {
	// 若所有 items 都對應同一 model，則建立葉節點
	firstModel := items[0].Model
	allSame := true
	for _, item := range items {
		if item.Model != firstModel {
			allSame = false
			break
		}
	}
	if allSame {
		return &Node{Model: firstModel}
	}

	// 搜集目前尚未使用過的 predicate 鍵
	candidateKeys := map[string]bool{}
	for _, item := range items {
		for k := range item.Predicate {
			if !usedKeys[k] {
				candidateKeys[k] = true
			}
		}
	}
	// 若已無候選鍵，則直接回傳第一個 model 作為 fallback
	if len(candidateKeys) == 0 {
		return &Node{Model: firstModel}
	}

	// 依照優先順序選取 splitKey
	var splitKey string
	// 先檢查優先鍵列表中是否有符合的
	for _, key := range priorityKeys {
		if candidateKeys[key] {
			splitKey = key
			break
		}
	}
	// 若沒有符合的優先鍵，隨機選取一個候選鍵
	if len(usedKeys) == 0 {
		if candidateKeys["custom_model_data"] {
			splitKey = "custom_model_data"
		}
	}

	// 如果 splitKey 還沒被設定，就根據優先鍵列表選擇
	if splitKey == "" {
		for _, key := range priorityKeys {
			if candidateKeys[key] {
				splitKey = key
				break
			}
		}
	}
	// 若依然沒選到，就隨機選取一個候選鍵
	if splitKey == "" {
		for k := range candidateKeys {
			splitKey = k
			break
		}
	}
	const missingValue = 0
	// 建立當前節點（此節點只是用來分支，不對應 model）
	root := &Node{}

	// 複製 usedKeys，並記錄 splitKey 已使用
	newUsedKeys := make(map[string]bool)
	for k, v := range usedKeys {
		newUsedKeys[k] = v
	}
	newUsedKeys[splitKey] = true

	// 根據 splitKey 的值將 items 分組
	groups := make(map[float64][]Item)
	// 如果某個項目沒有此鍵，則以一個特殊值（例如 -1）分組

	for _, item := range items {
		if val, ok := item.Predicate[splitKey]; ok {
			groups[val] = append(groups[val], item)
		} else {
			groups[missingValue] = append(groups[missingValue], item)
		}
	}

	// 針對每個群組建立子樹，並記錄該子樹的分割依據
	for val, groupItems := range groups {
		child := buildTree(groupItems, newUsedKeys)
		pair := PredicatePair{Key: splitKey, Value: val}
		child.Predicate = &pair
		if child.Predicate == nil {
			continue
		}
		key := child.Predicate.Key
		if root.SubNodes == nil {
			root.SubNodes = make(map[string][]*Node)
		}
		root.SubNodes[key] = append(root.SubNodes[key], child)
	}

	return root
}
func getModelByPath(param ModelParam) item_model.Model {
	modelName := param.ModelOverride
	if param.ModelOverride == "item/"+param.ModelBase {
		if param.Parent != "minecraft:item/generated" {
			modelName = param.Parent
		}
	}
	if param.ModelBase == "player_head" && (param.Parent == "minecraft:item/template_skull" || param.Parent == "item/template_skull") {
		return &special.Special{
			Model: &special.Head{
				Kind: special.HeadKindPlayer,
			},
			Base: modelName,
		}
	}

	return &model.Model{
		Model: modelName,
	}
}
