#pragma once
#include <string>
#include <vector>
#include <map>
#include "array.h"
#include "forward_list.h"
#include "list.h"
#include "stack.h"
#include "queue.h"
#include "tree.h"

using namespace std;

struct DB {
    map<string, DynArray<string>> arrs;
    map<string, SinglyList<string>> slists;
    map<string, DoublyList<string>> dlists;
    map<string, LifoStack<string>> stks;
    map<string, FifoQueue<string>> ques;
    map<string, CBTree<int>> trs;
};

vector<string> splitTokens(const string& s);
bool isNumber(const string& s);
string runQuery(DB& db, const string& line);
void loadFile(DB& db, const string& path);
void appendToFile(const string& path, const string& cmd);
