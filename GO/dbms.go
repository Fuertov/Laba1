package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DB struct {
	Arrs map[string]*DynArray[string]
	SLists map[string]*SinglyList[string]
	DLists  map[string]*DoublyList[string]
	Stks map[string]*LifoStack[string]
	Ques map[string]*FifoQueue[string]
	Trs  map[string]*CBTree[int]
}

func NewDB() *DB {
	return &DB{
		Arrs: make(map[string]*DynArray[string]),
		SLists: make(map[string]*SinglyList[string]),
		DLists:  make(map[string]*DoublyList[string]),
		Stks: make(map[string]*LifoStack[string]),
		Ques: make(map[string]*FifoQueue[string]),
		Trs:  make(map[string]*CBTree[int]),
	}
}

func splitTokens(s string) []string { return strings.Fields(s) }

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func runQuery(db *DB, line string) string {
	t := splitTokens(line)
	if len(t) == 0 || len(t[0]) == 0 {
		return "ОШИБКА: пустая команда"
	}
	
	op := t[0]
	
	if op == "PRINT" {
		if len(t) != 2 {
			return "ОШИБКА: PRINT <name>"
		}
		name := t[1]
		if db.Arrs[name] != nil {
			db.Arrs[name].MShow()
			return "OK"
		}
		if db.SLists[name] != nil {
			db.SLists[name].FShow()
			return "OK"
		}
		if db.DLists[name] != nil {
			db.DLists[name].LShow()
			return "OK"
		}
		if db.Stks[name] != nil {
			db.Stks[name].SShow()
			return "OK"
		}
		if db.Ques[name] != nil {
			db.Ques[name].QShow()
			return "OK"
		}
		if db.Trs[name] != nil {
			db.Trs[name].TShow()
			return "OK"
		}
		return "ОШИБКА: контейнер '" + name + "' не найден"
	}
	
	switch op[0] {
	case 'M':
		if op == "MPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: MPUSH <name> <value>"
			}
			if db.Arrs[t[1]] == nil {
				db.Arrs[t[1]] = &DynArray[string]{}
				db.Arrs[t[1]].Init()
			}
			db.Arrs[t[1]].MAppend(t[2])
			return t[2]
		}
		if op == "MINSERT" {
			if len(t) != 4 {
				return "ОШИБКА: MINSERT <name> <idx> <value>"
			}
			if !isNumber(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrs[t[1]] == nil {
				db.Arrs[t[1]] = &DynArray[string]{}
				db.Arrs[t[1]].Init()
			}
			idx, _ := strconv.Atoi(t[2])
			db.Arrs[t[1]].MInsertAt(idx, t[3])
			return t[3]
		}
		if op == "MGET" {
			if len(t) != 3 {
				return "ОШИБКА: MGET <name> <idx>"
			}
			if !isNumber(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrs[t[1]] == nil {
				return ""
			}
			idx, _ := strconv.Atoi(t[2])
			return db.Arrs[t[1]].MGet(idx)
		}
		if op == "MDEL" {
			if len(t) != 3 {
				return "ОШИБКА: MDEL <name> <idx>"
			}
			if !isNumber(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrs[t[1]] != nil {
				idx, _ := strconv.Atoi(t[2])
				db.Arrs[t[1]].MRemoveAt(idx)
			}
			return "OK"
		}
		if op == "MREPLACE" {
			if len(t) != 4 {
				return "ОШИБКА: MREPLACE <name> <idx> <value>"
			}
			if !isNumber(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrs[t[1]] == nil {
				db.Arrs[t[1]] = &DynArray[string]{}
				db.Arrs[t[1]].Init()
			}
			idx, _ := strconv.Atoi(t[2])
			db.Arrs[t[1]].MReplaceAt(idx, t[3])
			return t[3]
		}
		if op == "MSIZE" {
			if len(t) != 2 {
				return "ОШИБКА: MSIZE <name>"
			}
			if db.Arrs[t[1]] == nil {
				return "0"
			}
			return strconv.Itoa(db.Arrs[t[1]].MSize())
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'F':
		if op == "FPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: FPUSH <name> <value>"
			}
			if db.SLists[t[1]] == nil {
				db.SLists[t[1]] = &SinglyList[string]{}
			}
			db.SLists[t[1]].FAddHead(t[2])
			return t[2]
		}
		if op == "FPUSHEND" {
			if len(t) != 3 {
				return "ОШИБКА: FPUSHEND <name> <value>"
			}
			if db.SLists[t[1]] == nil {
				db.SLists[t[1]] = &SinglyList[string]{}
			}
			db.SLists[t[1]].FAddTail(t[2])
			return t[2]
		}
				if op == "FINSAFTER" {
			if len(t) != 4 {
				return "ОШИБКА: FINSAFTER <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *SNode[string]
			for c := db.SLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.SLists[t[1]].FAddAfter(node, t[3])
			return t[3]
		}
		if op == "FINSBEFORE" {
			if len(t) != 4 {
				return "ОШИБКА: FINSBEFORE <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *SNode[string]
			for c := db.SLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.SLists[t[1]].FAddBefore(node, t[3])
			return t[3]
		}
		if op == "FREMAFTER" {
			if len(t) != 3 {
				return "ОШИБКА: FREMAFTER <name> <ref>"
			}
			var node *SNode[string]
			for c := db.SLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.SLists[t[1]].FRemoveAfter(node)
			return "OK"
		}
		if op == "FREMBEFORE" {
			if len(t) != 3 {
				return "ОШИБКА: FREMBEFORE <name> <ref>"
			}
			var node *SNode[string]
			for c := db.SLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.SLists[t[1]].FRemoveBefore(node)
			return "OK"
		}
		if op == "FDEL" {
			if len(t) != 2 {
				return "ОШИБКА: FDEL <name>"
			}
			if db.SLists[t[1]] != nil {
				db.SLists[t[1]].FRemoveHead()
			}
			return "OK"
		}
		if op == "FDELEND" {
			if len(t) != 2 {
				return "ОШИБКА: FDELEND <name>"
			}
			if db.SLists[t[1]] != nil {
				db.SLists[t[1]].FRemoveTail()
			}
			return "OK"
		}
		if op == "FREMVAL" {
			if len(t) != 3 {
				return "ОШИБКА: FREMVAL <name> <value>"
			}
			if db.SLists[t[1]] != nil {
				db.SLists[t[1]].FRemoveValue(t[2])
			}
			return "OK"
		}
		if op == "FSEARCH" {
			if len(t) != 3 {
				return "ОШИБКА: FSEARCH <name> <value>"
			}
			if db.SLists[t[1]] == nil {
				fmt.Println("Элемент не был найден в списке")
			} else {
				db.SLists[t[1]].FSearch(t[2])
			}
			return "OK"
		}
		if op == "FSHOW" {
			if len(t) != 2 {
				return "ОШИБКА: FSHOW <name>"
			}
			if db.SLists[t[1]] == nil {
				fmt.Println("F[]")
			} else {
				db.SLists[t[1]].FShow()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'L':
		if op == "LPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: LPUSH <name> <value>"
			}
			if db.DLists[t[1]] == nil {
				db.DLists[t[1]] = &DoublyList[string]{}
			}
			db.DLists[t[1]].LAddHead(t[2])
			return t[2]
		}
		if op == "LPUSHEND" {
			if len(t) != 3 {
				return "ОШИБКА: LPUSHEND <name> <value>"
			}
			if db.DLists[t[1]] == nil {
				db.DLists[t[1]] = &DoublyList[string]{}
			}
			db.DLists[t[1]].LAddTail(t[2])
			return t[2]
		}
				if op == "LINSAFTER" {
			if len(t) != 4 {
				return "ОШИБКА: LINSAFTER <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *DNode[string]
			for c := db.DLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.DLists[t[1]].LAddAfter(node, t[3])
			return t[3]
		}
		if op == "LINSBEFORE" {
			if len(t) != 4 {
				return "ОШИБКА: LINSBEFORE <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *DNode[string]
			for c := db.DLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.DLists[t[1]].LAddBefore(node, t[3])
			return t[3]
		}
		if op == "LREMAFTER" {
			if len(t) != 3 {
				return "ОШИБКА: LREMAFTER <name> <ref>"
			}
			var node *DNode[string]
			for c := db.DLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.DLists[t[1]].LRemoveAfter(node)
			return "OK"
		}
		if op == "LREMBEFORE" {
			if len(t) != 3 {
				return "ОШИБКА: LREMBEFORE <name> <ref>"
			}
			var node *DNode[string]
			for c := db.DLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.DLists[t[1]].LRemoveBefore(node)
			return "OK"
		}
		if op == "LDEL" {
			if len(t) != 2 {
				return "ОШИБКА: LDEL <name>"
			}
			if db.DLists[t[1]] != nil {
				db.DLists[t[1]].LRemoveHead()
			}
			return "OK"
		}
		if op == "LDELEND" {
			if len(t) != 2 {
				return "ОШИБКА: LDELEND <name>"
			}
			if db.DLists[t[1]] != nil {
				db.DLists[t[1]].LRemoveTail()
			}
			return "OK"
		}
		if op == "LREMVAL" {
			if len(t) != 3 {
				return "ОШИБКА: LREMVAL <name> <value>"
			}
			if db.DLists[t[1]] != nil {
				db.DLists[t[1]].LRemoveValue(t[2])
			}
			return "OK"
		}
		if op == "LSEARCH" {
			if len(t) != 3 {
				return "ОШИБКА: LSEARCH <name> <value>"
			}
			if db.DLists[t[1]] == nil {
				fmt.Println("Элемент не был найден в списке")
			} else {
				db.DLists[t[1]].LSearch(t[2])
			}
			return "OK"
		}
		if op == "LSHOW" {
			if len(t) != 2 {
				return "ОШИБКА: LSHOW <name>"
			}
			if db.DLists[t[1]] == nil {
				fmt.Println("L[]")
			} else {
				db.DLists[t[1]].LShow()
			}
			return "OK"
		}
		if op == "LSHOWREV" {
			if len(t) != 2 {
				return "ОШИБКА: LSHOWREV <name>"
			}
			if db.DLists[t[1]] == nil {
				fmt.Println("Lr[]")
			} else {
				db.DLists[t[1]].LShowReverse()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'S':
		if op == "SPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: SPUSH <name> <value>"
			}
			if db.Stks[t[1]] == nil {
				db.Stks[t[1]] = &LifoStack[string]{}
			}
			db.Stks[t[1]].SAdd(t[2])
			return t[2]
		}
		if op == "SPOP" {
			if len(t) != 2 {
				return "ОШИБКА: SPOP <name>"
			}
			if db.Stks[t[1]] == nil || db.Stks[t[1]].SIsEmpty() {
				return "ОШИБКА: стек пуст"
			}
			v := db.Stks[t[1]].SPeek()
			db.Stks[t[1]].SRemove()
			return v
		}
		if op == "SPEEK" {
			if len(t) != 2 {
				return "ОШИБКА: SPEEK <name>"
			}
			if db.Stks[t[1]] == nil || db.Stks[t[1]].SIsEmpty() {
				return "ОШИБКА: стек пуст"
			}
			return db.Stks[t[1]].SPeek()
		}
		if op == "SSHOW" {
			if len(t) != 2 {
				return "ОШИБКА: SSHOW <name>"
			}
			if db.Stks[t[1]] == nil {
				fmt.Println("S[]")
			} else {
				db.Stks[t[1]].SShow()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'Q':
		if op == "QPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: QPUSH <name> <value>"
			}
			if db.Ques[t[1]] == nil {
				db.Ques[t[1]] = &FifoQueue[string]{}
			}
			db.Ques[t[1]].QAdd(t[2])
			return t[2]
		}
		if op == "QPOP" {
			if len(t) != 2 {
				return "ОШИБКА: QPOP <name>"
			}
			if db.Ques[t[1]] == nil || db.Ques[t[1]].QIsEmpty() {
				return "ОШИБКА: очередь пуста"
			}
			v := db.Ques[t[1]].QPeek()
			db.Ques[t[1]].QRemove()
			return v
		}
		if op == "QPEEK" {
			if len(t) != 2 {
				return "ОШИБКА: QPEEK <name>"
			}
			if db.Ques[t[1]] == nil || db.Ques[t[1]].QIsEmpty() {
				return "ОШИБКА: очередь пуста"
			}
			return db.Ques[t[1]].QPeek()
		}
		if op == "QSHOW" {
			if len(t) != 2 {
				return "ОШИБКА: QSHOW <name>"
			}
			if db.Ques[t[1]] == nil {
				fmt.Println("Q[]")
			} else {
				db.Ques[t[1]].QShow()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'T':
		if op == "TINSERT" {
			if len(t) != 3 {
				return "ОШИБКА: TINSERT <name> <int>"
			}
			if !isNumber(t[2]) {
				return "ОШИБКА: значение — не число"
			}
			if db.Trs[t[1]] == nil {
				db.Trs[t[1]] = &CBTree[int]{}
			}
			v, _ := strconv.Atoi(t[2])
			db.Trs[t[1]].TAdd(v)
			return t[2]
		}
		if op == "TSEARCH" {
			if len(t) != 3 {
				return "ОШИБКА: TSEARCH <name> <int>"
			}
			if !isNumber(t[2]) {
				return "ОШИБКА: значение — не число"
			}
			if db.Trs[t[1]] != nil {
				v, _ := strconv.Atoi(t[2])
				db.Trs[t[1]].TSearch(v)
			}
			return "OK"
		}
		if op == "TSHOW" {
			if len(t) != 2 {
				return "ОШИБКА: TSHOW <name>"
			}
			if db.Trs[t[1]] == nil {
				fmt.Println("T[]")
			} else {
				db.Trs[t[1]].TShow()
			}
			return "OK"
		}
		if op == "TSHOWTREE" {
			if len(t) != 2 {
				return "ОШИБКА: TSHOWTREE <name>"
			}
			if db.Trs[t[1]] != nil {
				db.Trs[t[1]].TShowTree()
			}
			return "OK"
		}
		if op == "TCHECKCOMPLETE" {
			if len(t) != 2 {
				return "ОШИБКА: TCHECKCOMPLETE <name>"
			}
			if db.Trs[t[1]] != nil && db.Trs[t[1]].TCheckComplete() {
				return "TRUE"
			}
			return "FALSE"
		}
		if op == "TCHECKFULL" {
			if len(t) != 2 {
				return "ОШИБКА: TCHECKFULL <name>"
			}
			if db.Trs[t[1]] != nil && db.Trs[t[1]].TCheckFull() {
				return "TRUE"
			}
			return "FALSE"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	default:
		return "ОШИБКА: неизвестная команда '" + op + "'"
	}
}
