#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct TNode {
    T value;
    TNode<T>* left = nullptr;
    TNode<T>* right = nullptr;
};

template <typename T>
struct CBTree {
    TNode<T>* root = nullptr;
    int size = 0;

    void TAdd(const T& v) {
        root = placeAt(root, size, 0, v);
        ++size;
    }

    TNode<T>* placeAt(TNode<T>* r, int target, int idx, const T& v) {
        if (!r) {
            if (idx == target)
                return new TNode<T>{ v, nullptr, nullptr };
            return nullptr;
        }
        if (2 * idx + 1 <= target)
            r->left = placeAt(r->left, target, 2 * idx + 1, v);
        if (2 * idx + 2 <= target)
            r->right = placeAt(r->right, target, 2 * idx + 2, v);
        return r;
    }

    void TSearch(const T& v) const {
        int depth = 0;
        if (searchRec(root, v, depth))
            cout << "Элемент был найден на глубине " << depth << endl;
        else
            cout << "Элемент не был найден в дереве" << endl;
    }

    bool searchRec(TNode<T>* r, const T& v, int& depth) const {
        if (!r)
            return false;
        else if (r->value == v)
            return true;
        else if (searchRec(r->left, v, depth)) {
            ++depth;
            return true;
        }
        else if (searchRec(r->right, v, depth)) {
            ++depth;
            return true;
        }
        return false;
    }

    void TShow() const {
        cout << "T[";
        inOrder(root);
        cout << "]" << endl;
    }

    void inOrder(TNode<T>* r) const {
        if (!r)
            return;
        inOrder(r->left);
        cout << r->value << " ";
        inOrder(r->right);
    }

    void TShowTree() const {
        drawTree(root, 0);
    }

    void drawTree(TNode<T>* r, int level) const {
        if (!r)
            return;
        drawTree(r->right, level + 1);
        for (int i = 0; i < level; ++i)
            cout << "\t";
        cout << r->value << endl;
        drawTree(r->left, level + 1);
    }

    bool TCheckComplete() const {
        return checkCompleteRec(root, 0, size);
    }

    bool checkCompleteRec(TNode<T>* r, int i, int n) const {
        if (!r)
            return true;
        if (i >= n)
            return false;
        return checkCompleteRec(r->left, 2 * i + 1, n)
            && checkCompleteRec(r->right, 2 * i + 2, n);
    }

    bool TCheckFull() const {
        return checkFullRec(root);
    }

    bool checkFullRec(TNode<T>* r) const {
        if (!r)
            return true;
        if (!r->left && !r->right)
            return true;
        if (r->left && r->right)
            return checkFullRec(r->left) && checkFullRec(r->right);
        return false;
    }

    void clear() {
        freeTree(root);
        root = nullptr;
        size = 0;
    }

    void freeTree(TNode<T>* r) {
        if (!r)
            return;
        freeTree(r->left);
        freeTree(r->right);
        delete r;
    }
};
