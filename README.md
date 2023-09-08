# Curiously recurring template pattern
This is a pattern in C++ where a base class has its derived class as
a template parameter:
```C++
template<typename T>
struct Base {
    void speak() {
        std::cout << T::message() << std::endl;
    }
};

struct Derived : public Base<Derived> {
    static std::string message() { return "Hello!"; }
};

int main() {
    Derived{}.speak(); // -> Prints 'Hello!'
}
```

I wondered whether something similar would work in Go. Sure enough, it works!
This repo shows an example.