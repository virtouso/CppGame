#pragma once
#include <string>

class component
{
public:
    virtual ~component() = default;
    virtual std::string GetName() =0;
};
