#pragma once
#include <iostream>
#include "forward_list.h"

using namespace std;

template <typename T>
struct LifoStack {
    SNode<T>* head = nullptr;

    bool SIsEmpty() const {
        return head == nullptr;
    }

    void SAdd(const T& v) {
        head = new SNode<T>{ v, head };
    }

    void SRemove() {
        if (SIsEmpty()) {
            cout << "ОШИБКА: Стек пуст, удалять нечего" << endl;
            return;
        }
        SNode<T>* d = head;
        head = d->next;
        delete d;
    }

    T SPeek() const {
        if (SIsEmpty()) {
            cout << "ОШИБКА: Стек пуст" << endl;
            return T{};
        }
        return head->value;
    }

    void SShow() const {
        cout << "S[";
        for (SNode<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void clear() {
        while (head) {
            SNode<T>* n = head->next;
            delete head;
            head = n;
        }
    }
};
