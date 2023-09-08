#include <iostream>

template <typename T>
struct Base {
    void speak() { std::cout << T::message() << std::endl; }
};

struct Derived : public Base<Derived> {
    static constexpr std::string_view message() { return "Hello!"; }
};

int main() {
    Derived{}.speak();  // -> Prints Hello!
}