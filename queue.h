#pragma once
#include <iostream>
#include "forward_list.h"

using namespace std;

template <typename T>
struct FifoQueue {
    SNode<T>* head = nullptr;
    SNode<T>* tail = nullptr;

    bool QIsEmpty() const {
        return head == nullptr;
    }

    void QAdd(const T& v) {
        SNode<T>* n = new SNode<T>{ v, nullptr };
        if (!tail) {
            head = n;
            tail = n;
        }
        else {
            tail->next = n;
            tail = n;
        }
    }

    void QRemove() {
        if (QIsEmpty()) {
            cout << "ОШИБКА: Очередь пуста, удалять нечего" << endl;
            return;
        }
        SNode<T>* d = head;
        head = d->next;
        delete d;
        if (!head)
            tail = nullptr;
    }

    T QPeek() const {
        if (QIsEmpty()) {
            cout << "ОШИБКА: Очередь пуста" << endl;
            return T{};
        }
        return head->value;
    }

    void QShow() const {
        cout << "Q[";
        for (SNode<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void clear() {
        SNode<T>* c = head;
        while (c) {
            SNode<T>* n = c->next;
            delete c;
            c = n;
        }
        head = nullptr;
        tail = nullptr;
    }
};
