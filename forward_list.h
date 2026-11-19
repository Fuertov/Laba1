#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct SNode {
    T value;
    SNode<T>* next = nullptr;
};

template <typename T>
struct SinglyList {
    SNode<T>* head = nullptr;

    void FAddHead(const T& v) {
        head = new SNode<T>{ v, head };
    }

    void FAddTail(const T& v) {
        SNode<T>* n = new SNode<T>{ v, nullptr };
        if (!head) {
            head = n;
            return;
        }
        SNode<T>* cur = head;
        while (cur->next)
            cur = cur->next;
        cur->next = n;
    }

    void FAddAfter(SNode<T>* node, const T& v) {
        if (node == nullptr) {
            FAddTail(v);
            return;
        }
        SNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (cur)
            cur->next = new SNode<T>{ v, cur->next };
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void FAddBefore(SNode<T>* node, const T& v) {
        if (head == node || node == nullptr) {
            FAddHead(v);
            return;
        }
        SNode<T>* cur = head;
        while (cur && cur->next != node)
            cur = cur->next;
        if (cur)
            cur->next = new SNode<T>{ v, cur->next };
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void FRemoveHead() {
        if (!head) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        SNode<T>* d = head;
        head = d->next;
        delete d;
    }

    void FRemoveTail() {
        if (!head) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        else if (!head->next) {
            delete head;
            head = nullptr;
            return;
        }
        SNode<T>* cur = head;
        while (cur->next->next)
            cur = cur->next;
        delete cur->next;
        cur->next = nullptr;
    }

    void FRemoveAfter(SNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        SNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (cur) {
            if (cur->next == nullptr) {
                cout << "ОШИБКА: После узла " << node->value << " ничего нет" << endl;
                return;
            }
            SNode<T>* d = cur->next;
            cur->next = d->next;
            delete d;
        }
        else {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
        }
    }

    void FRemoveBefore(SNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (head == node) {
            cout << "ОШИБКА: Перед головой ничего нет" << endl;
            return;
        }
        else if (head->next == node) {
            FRemoveHead();
            return;
        }
        SNode<T>* cur = head;
        while (cur && cur->next && cur->next->next != node)
            cur = cur->next;
        if (cur) {
            SNode<T>* d = cur->next;
            cur->next = d->next;
            delete d;
        }
        else {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
        }
    }

    void FRemoveNode(SNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (head == node) {
            FRemoveHead();
            return;
        }
        SNode<T>* cur = head;
        while (cur && cur->next != node)
            cur = cur->next;
        if (cur) {
            cur->next = node->next;
            delete node;
        }
        else {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
        }
    }

    void FRemoveValue(const T& v) {
        SNode<T>* cur = head;
        while (cur && cur->value != v)
            cur = cur->next;
        if (cur)
            FRemoveNode(cur);
        else
            cout << "ОШИБКА: Значения " << v << " нет в списке" << endl;
    }

    void FSearch(const T& v) {
        int i = 0;
        for (SNode<T>* c = head; c; c = c->next, ++i) {
            if (c->value == v) {
                cout << "Элемент был найден на позиции " << i << endl;
                return;
            }
        }
        cout << "Элемент не был найден в списке" << endl;
    }

    void FShow() {
        cout << "F[";
        for (SNode<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void FShowRec(SNode<T>* n) {
        if (!n)
            return;
        cout << n->value << " ";
        FShowRec(n->next);
    }

    void clear() {
        while (head) {
            SNode<T>* n = head->next;
            delete head;
            head = n;
        }
    }
};
