#pragma once
#include <vector>

#include "component.h"

class game_object
{
public:
    std::vector<component*> components;


    ~game_object()
    {
        for (auto obj : components)
        {
            delete obj;
        }
        components.clear();
        components.shrink_to_fit();
    }
};
