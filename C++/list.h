#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct DNode {
    T value;
    DNode<T>* next = nullptr;
    DNode<T>* prev = nullptr;
};

template <typename T>
struct DoublyList {
    DNode<T>* head = nullptr;
    DNode<T>* tail = nullptr;

    void LAddHead(const T& v) {
        DNode<T>* n = new DNode<T>{ v, head, nullptr };
        if (head)
            head->prev = n;
        else
            tail = n;
        head = n;
    }

    void LAddTail(const T& v) {
        DNode<T>* n = new DNode<T>{ v, nullptr, tail };
        if (tail)
            tail->next = n;
        else
            head = n;
        tail = n;
    }

    void LAddBefore(DNode<T>* node, const T& v) {
        if (node == nullptr || head == node) {
            LAddHead(v);
            return;
        }
        DNode<T>* cur = head;
        while (cur && cur->next != node)
            cur = cur->next;
        if (cur) {
            DNode<T>* n = new DNode<T>{ v, node, node->prev };
            node->prev->next = n;
            node->prev = n;
        }
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void LAddAfter(DNode<T>* node, const T& v) {
        if (node == nullptr) {
            LAddTail(v);
            return;
        }
        DNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (cur) {
            DNode<T>* n = new DNode<T>{ v, node->next, node };
            if (node->next)
                node->next->prev = n;
            else
                tail = n;
            node->next = n;
        }
        else
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
    }

    void LRemoveHead() {
        if (!head) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        DNode<T>* d = head;
        head = d->next;
        if (head)
            head->prev = nullptr;
        else
            tail = nullptr;
        delete d;
    }

    void LRemoveTail() {
        if (!tail) {
            cout << "ОШИБКА: В списке нет элементов, удалять нечего" << endl;
            return;
        }
        DNode<T>* d = tail;
        tail = d->prev;
        if (tail)
            tail->next = nullptr;
        else
            head = nullptr;
        delete d;
    }

    void LRemoveNode(DNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        DNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (!cur) {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
            return;
        }
        if (node->prev)
            node->prev->next = node->next;
        else
            head = node->next;
        if (node->next)
            node->next->prev = node->prev;
        else
            tail = node->prev;
        delete node;
    }

    void LRemoveAfter(DNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (node == tail) {
            cout << "ОШИБКА: За последним элементом, нет элементов для удаления" << endl;
            return;
        }
        DNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (!cur) {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
            return;
        }
        else if (cur->next == tail)
            LRemoveTail();
        else {
            DNode<T>* d = cur->next;
            cur->next = d->next;
            cur->next->prev = d->prev;
            delete d;
        }
    }

    void LRemoveBefore(DNode<T>* node) {
        if (node == nullptr) {
            cout << "ОШИБКА: Узел не может быть nullptr" << endl;
            return;
        }
        else if (node == head) {
            cout << "ОШИБКА: Перед первым элементом, нет элементов для удаления" << endl;
            return;
        }
        DNode<T>* cur = head;
        while (cur && cur != node)
            cur = cur->next;
        if (!cur) {
            cout << "ОШИБКА: Узла " << node->value << " нет в списке" << endl;
            return;
        }
        else if (cur->prev == head)
            LRemoveHead();
        else {
            DNode<T>* d = cur->prev;
            cur->prev = d->prev;
            cur->prev->next = d->next;
            delete d;
        }
    }

    void LRemoveValue(const T& v) {
        DNode<T>* cur = head;
        while (cur && cur->value != v)
            cur = cur->next;
        if (cur)
            LRemoveNode(cur);
        else
            cout << "ОШИБКА: Значения " << v << " нет в списке" << endl;
    }

    void LSearch(const T& v) {
        int i = 0;
        for (DNode<T>* c = head; c; c = c->next, ++i) {
            if (c->value == v) {
                cout << "Элемент был найден на позиции " << i << endl;
                return;
            }
        }
        cout << "Элемент не был найден в списке" << endl;
    }

    void LShow() {
        cout << "L[";
        for (DNode<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void LShowReverse() {
        cout << "Lr[";
        for (DNode<T>* c = tail; c; c = c->prev)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void LShowRec(DNode<T>* n) {
        if (!n)
            return;
        cout << n->value << " ";
        LShowRec(n->next);
    }

    void clear() {
        DNode<T>* c = head;
        while (c) {
            DNode<T>* n = c->next;
            delete c;
            c = n;
        }
        head = nullptr;
        tail = nullptr;
    }
};
