#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct DynArray {
    T* data = nullptr;
    int size = 0;
    int capacity = 0;

    void Init() {
        capacity = 1;
        size = 0;
        data = new T[capacity];
    }

    void MAppend(const T& n) {
        if (size == capacity) {
            int newCap = capacity ? capacity * 2 : 1;
            T* nd = new T[newCap];
            for (int i = 0; i < size; ++i)
                nd[i] = data[i];
            delete[] data;
            data = nd;
            capacity = newCap;
        }
        data[size++] = n;
    }

    void MInsertAt(int idx, const T& n) {
        if (idx < 0 || idx > size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return;
        }
        MAppend(T{});
        for (int i = size - 1; i > idx; --i)
            data[i] = data[i - 1];
        data[idx] = n;
    }

    T MGet(int idx) const {
        if (idx < 0 || idx >= size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return T{};
        }
        return data[idx];
    }

    void MRemoveAt(int idx) {
        if (idx < 0 || idx >= size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return;
        }
        for (int i = idx; i < size - 1; ++i)
            data[i] = data[i + 1];
        size--;
    }

    void MReplaceAt(int idx, const T& v) {
        if (idx < 0 || idx >= size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return;
        }
        data[idx] = v;
    }

    int MSize() const { return size; }

    void MShow() const {
        cout << "M[";
        for (int i = 0; i < size; i++)
            cout << data[i] << " ";
        cout << "]" << endl;
    }

    void clear() {
        delete[] data;
        data = nullptr;
        size = 0;
        capacity = 0;
    }
};
