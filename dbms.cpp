#include "dbms.h"
#include <iostream>
#include <fstream>

using namespace std;

vector<string> splitTokens(const string& s) {
    vector<string> out;
    string cur;
    for (char c : s) {
        if (c == ' ') {
            if (!cur.empty()) {
                out.push_back(cur);
                cur.clear();
            }
        }
        else
            cur.push_back(c);
    }
    if (!cur.empty())
        out.push_back(cur);
    return out;
}

bool isNumber(const string& s) {
    if (s.empty())
        return false;
    int i = 0;
    if (s[0] == '-') {
        if (s.size() == 1)
            return false;
        i = 1;
    }
    for (; i < (int)s.size(); ++i)
        if (s[i] < '0' || s[i] > '9')
            return false;
    return true;
}

string runQuery(DB& db, const string& line) {
    auto t = splitTokens(line);
    if (t.empty())
        return "ОШИБКА: пустая команда";
    
    string op = t[0];
	
    if (op == "PRINT") {
        if (t.size() != 2) return "ОШИБКА: PRINT <name>";
        string name = t[1];
        if (db.arrs.count(name)) { db.arrs[name].MShow(); return "OK"; }
        if (db.slists.count(name)) { db.slists[name].FShow(); return "OK"; }
        if (db.dlists.count(name))  { db.dlists[name].LShow();  return "OK"; }
        if (db.stks.count(name)) { db.stks[name].SShow(); return "OK"; }
        if (db.ques.count(name)) { db.ques[name].QShow(); return "OK"; }
        if (db.trs.count(name))  { db.trs[name].TShow();  return "OK"; }
        return "ОШИБКА: контейнер '" + name + "' не найден";
    }
    
    switch (op[0]) {
    case 'M': {
        if (op == "MPUSH") {
            if (t.size() != 3) return "ОШИБКА: MPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.arrs[t[1]].MAppend(t[2]);
            return t[2];
        }
        if (op == "MINSERT") {
            if (t.size() != 4) return "ОШИБКА: MINSERT <name> <idx> <value>";
            if (!isNumber(t[2]))  return "ОШИБКА: idx — не число";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            db.arrs[t[1]].MInsertAt(stoi(t[2]), t[3]);
            return t[3];
        }
        if (op == "MGET") {
            if (t.size() != 3) return "ОШИБКА: MGET <name> <idx>";
            if (!isNumber(t[2]))  return "ОШИБКА: idx — не число";
            return db.arrs[t[1]].MGet(stoi(t[2]));
        }
        if (op == "MDEL") {
            if (t.size() != 3) return "ОШИБКА: MDEL <name> <idx>";
            if (!isNumber(t[2]))  return "ОШИБКА: idx — не число";
            db.arrs[t[1]].MRemoveAt(stoi(t[2]));
            return "OK";
        }
        if (op == "MREPLACE") {
            if (t.size() != 4) return "ОШИБКА: MREPLACE <name> <idx> <value>";
            if (!isNumber(t[2]))  return "ОШИБКА: idx — не число";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            db.arrs[t[1]].MReplaceAt(stoi(t[2]), t[3]);
            return t[3];
        }
        if (op == "MSIZE") {
            if (t.size() != 2) return "ОШИБКА: MSIZE <name>";
            return to_string(db.arrs[t[1]].MSize());
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'F': {
        if (op == "FPUSH") {
            if (t.size() != 3) return "ОШИБКА: FPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.slists[t[1]].FAddHead(t[2]);
            return t[2];
        }
        if (op == "FPUSHEND") {
            if (t.size() != 3) return "ОШИБКА: FPUSHEND <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.slists[t[1]].FAddTail(t[2]);
            return t[2];
        }
        if (op == "FDEL") {
            if (t.size() != 2) return "ОШИБКА: FDEL <name>";
            db.slists[t[1]].FRemoveHead();
            return "OK";
        }
        if (op == "FDELEND") {
            if (t.size() != 2) return "ОШИБКА: FDELEND <name>";
            db.slists[t[1]].FRemoveTail();
            return "OK";
        }
        if (op == "FINSAFTER") {
	    if (t.size() != 4) return "ОШИБКА: FINSAFTER <name> <ref> <value>";
	    SNode<string>* node = nullptr;
	    for (SNode<string>* c = db.slists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.slists[t[1]].FAddAfter(node, t[3]);
	    return t[3];
	}
		if (op == "FINSBEFORE") {
	    if (t.size() != 4) return "ОШИБКА: FINSBEFORE <name> <ref> <value>";
	    SNode<string>* node = nullptr;
	    for (SNode<string>* c = db.slists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.slists[t[1]].FAddBefore(node, t[3]);
	    return t[3];
	}
	if (op == "FREMAFTER") {
	    if (t.size() != 3) return "ОШИБКА: FREMAFTER <name> <ref>";
	    SNode<string>* node = nullptr;
	    for (SNode<string>* c = db.slists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.slists[t[1]].FRemoveAfter(node);
	    return "OK";
	}
	if (op == "FREMBEFORE") {
	    if (t.size() != 3) return "ОШИБКА: FREMBEFORE <name> <ref>";
	    SNode<string>* node = nullptr;
	    for (SNode<string>* c = db.slists[t[1]].head; c; c = c->next)
		if (c->value == t[2]) { node = c; break; }
	    if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
	    db.slists[t[1]].FRemoveBefore(node);
	    return "OK";
	}
        if (op == "FREMVAL") {
            if (t.size() != 3) return "ОШИБКА: FREMVAL <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.slists[t[1]].FRemoveValue(t[2]);
            return "OK";
        }
        if (op == "FSEARCH") {
            if (t.size() != 3) return "ОШИБКА: FSEARCH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.slists[t[1]].FSearch(t[2]);
            return "OK";
        }
        if (op == "FSHOW") {
            if (t.size() != 2) return "ОШИБКА: FSHOW <name>";
            db.slists[t[1]].FShow();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'L': {
        if (op == "LPUSH") {
            if (t.size() != 3) return "ОШИБКА: LPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.dlists[t[1]].LAddHead(t[2]);
            return t[2];
        }
        if (op == "LPUSHEND") {
            if (t.size() != 3) return "ОШИБКА: LPUSHEND <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.dlists[t[1]].LAddTail(t[2]);
            return t[2];
        }
                if (op == "LINSAFTER") {
            if (t.size() != 4) return "ОШИБКА: LINSAFTER <name> <ref> <value>";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            DNode<string>* node = nullptr;
            for (DNode<string>* c = db.dlists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.dlists[t[1]].LAddAfter(node, t[3]);
            return t[3];
        }
        if (op == "LINSBEFORE") {
            if (t.size() != 4) return "ОШИБКА: LINSBEFORE <name> <ref> <value>";
            if (t[3].empty())  return "ОШИБКА: пустое значение";
            DNode<string>* node = nullptr;
            for (DNode<string>* c = db.dlists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.dlists[t[1]].LAddBefore(node, t[3]);
            return t[3];
        }
        if (op == "LREMAFTER") {
            if (t.size() != 3) return "ОШИБКА: LREMAFTER <name> <ref>";
            DNode<string>* node = nullptr;
            for (DNode<string>* c = db.dlists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.dlists[t[1]].LRemoveAfter(node);
            return "OK";
        }
        if (op == "LREMBEFORE") {
            if (t.size() != 3) return "ОШИБКА: LREMBEFORE <name> <ref>";
            DNode<string>* node = nullptr;
            for (DNode<string>* c = db.dlists[t[1]].head; c; c = c->next)
                if (c->value == t[2]) { node = c; break; }
            if (!node) return "ОШИБКА: узла '" + t[2] + "' нет в списке";
            db.dlists[t[1]].LRemoveBefore(node);
            return "OK";
        }
        if (op == "LDEL") {
            if (t.size() != 2) return "ОШИБКА: LDEL <name>";
            db.dlists[t[1]].LRemoveHead();
            return "OK";
        }
        if (op == "LDELEND") {
            if (t.size() != 2) return "ОШИБКА: LDELEND <name>";
            db.dlists[t[1]].LRemoveTail();
            return "OK";
        }
        if (op == "LREMVAL") {
            if (t.size() != 3) return "ОШИБКА: LREMVAL <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.dlists[t[1]].LRemoveValue(t[2]);
            return "OK";
        }
        if (op == "LSEARCH") {
            if (t.size() != 3) return "ОШИБКА: LSEARCH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.dlists[t[1]].LSearch(t[2]);
            return "OK";
        }
        if (op == "LSHOW") {
            if (t.size() != 2) return "ОШИБКА: LSHOW <name>";
            db.dlists[t[1]].LShow();
            return "OK";
        }
        if (op == "LSHOWREV") {
            if (t.size() != 2) return "ОШИБКА: LSHOWREV <name>";
            db.dlists[t[1]].LShowReverse();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'S': {
        if (op == "SPUSH") {
            if (t.size() != 3) return "ОШИБКА: SPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.stks[t[1]].SAdd(t[2]);
            return t[2];
        }
        if (op == "SPOP") {
            if (t.size() != 2) return "ОШИБКА: SPOP <name>";
            if (db.stks[t[1]].SIsEmpty()) return "ОШИБКА: стек пуст";
            string v = db.stks[t[1]].SPeek();
            db.stks[t[1]].SRemove();
            return v;
        }
        if (op == "SPEEK") {
            if (t.size() != 2) return "ОШИБКА: SPEEK <name>";
            if (db.stks[t[1]].SIsEmpty()) return "ОШИБКА: стек пуст";
            return db.stks[t[1]].SPeek();
        }
        if (op == "SSHOW") {
            if (t.size() != 2) return "ОШИБКА: SSHOW <name>";
            db.stks[t[1]].SShow();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'Q': {
        if (op == "QPUSH") {
            if (t.size() != 3) return "ОШИБКА: QPUSH <name> <value>";
            if (t[2].empty())  return "ОШИБКА: пустое значение";
            db.ques[t[1]].QAdd(t[2]);
            return t[2];
        }
        if (op == "QPOP") {
            if (t.size() != 2) return "ОШИБКА: QPOP <name>";
            if (db.ques[t[1]].QIsEmpty()) return "ОШИБКА: очередь пуста";
            string v = db.ques[t[1]].QPeek();
            db.ques[t[1]].QRemove();
            return v;
        }
        if (op == "QPEEK") {
            if (t.size() != 2) return "ОШИБКА: QPEEK <name>";
            if (db.ques[t[1]].QIsEmpty()) return "ОШИБКА: очередь пуста";
            return db.ques[t[1]].QPeek();
        }
        if (op == "QSHOW") {
            if (t.size() != 2) return "ОШИБКА: QSHOW <name>";
            db.ques[t[1]].QShow();
            return "OK";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    case 'T': {
        if (op == "TINSERT") {
            if (t.size() != 3) return "ОШИБКА: TINSERT <name> <int>";
            if (!isNumber(t[2]))  return "ОШИБКА: значение — не число";
            db.trs[t[1]].TAdd(stoi(t[2]));
            return t[2];
        }
        if (op == "TSEARCH") {
            if (t.size() != 3) return "ОШИБКА: TSEARCH <name> <int>";
            if (!isNumber(t[2]))  return "ОШИБКА: значение — не число";
            db.trs[t[1]].TSearch(stoi(t[2]));
            return "OK";
        }
        if (op == "TSHOW") {
            if (t.size() != 2) return "ОШИБКА: TSHOW <name>";
            db.trs[t[1]].TShow();
            return "OK";
        }
        if (op == "TSHOWTREE") {
            if (t.size() != 2) return "ОШИБКА: TSHOWTREE <name>";
            db.trs[t[1]].TShowTree();
            return "OK";
        }
        if (op == "TCHECKCOMPLETE") {
            if (t.size() != 2) return "ОШИБКА: TCHECKCOMPLETE <name>";
            return db.trs[t[1]].TCheckComplete() ? "TRUE" : "FALSE";
        }
        if (op == "TCHECKFULL") {
            if (t.size() != 2) return "ОШИБКА: TCHECKFULL <name>";
            return db.trs[t[1]].TCheckFull() ? "TRUE" : "FALSE";
        }
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }

    default:
        return "ОШИБКА: неизвестная команда '" + op + "'";
    }
}

void loadFile(DB& db, const string& path) {
    ifstream in(path);
    string line;
    while (getline(in, line))
        if (!line.empty())
            runQuery(db, line);
}

void appendToFile(const string& path, const string& cmd) {
    ofstream out(path, ios::app);
    out << cmd << "\n";
}
