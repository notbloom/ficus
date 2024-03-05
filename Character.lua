local Deck = require "Deck"
local Card = require "Card"

local Character = {
    Name = "Bloom",
    Description = "Coder and Gamer",
    MaxCapacity = 10,
    StartingCapacity = 4,
    Capacity = 4,
    Energy = 2,
    Tags = {"Bloom", "Coder", "Gamer"},
    Weakness = {"Light", "Sound", "Social", "Unexpected"},
    Deck = "",
}

function Character:Sleep()

end
function Character:DeckIndexes()
    return {0,1,2,3,4,5,6,7,8,0,1,2,3,4,5,6,7,8,0,1,2,3,4,5,6,7,8}
end 
function Character:New (o)
    o = o or {}   -- create object if user does not provide one
    setmetatable(o, self)
    self.__index = self
    self.Deck = Deck:New({
        deck = {1,2,3,4,5,6,7,8,0,1,2,3,4,5,6,7,8,0,1,2,3,4,5,6,7,8}, -- self:DeckIndexes(),        
        repositoryMethod = self.GetCard,
    })
    return o
end

function Character:GetCard(index)
    if index == 0 then
        return Card:New({    
            Title = "Lentes de sol",
            Description = "Inmunidad a la luz",

        })
    elseif index == 1  then
        return Card:New({            
            Title = "Noche de Peliculas!",
            Description = "Termina el día",

        })
    elseif index == 2  then
        return Card:New({    
            Title = "Snack",
            Description = "Otorga 2 de Energia y Roba una carta",
        })
    elseif index == 3  then
        return Card:New({        
            Title = "Jugar Online",
            Description = "Roll: Están tus amigos online!, Roll: Terminar tu turno!"
        })
    elseif index == 4  then
        return Card:New({       
            Title = "Café",
            Description = "+2 Energía"
        })
    elseif index == 5  then
        return Card:New({       
            Title = "Ducha",
            Description = "+1 Energía y +1 Comodidad"
        })
    elseif index == 6  then
        return Card:New({       
            Title = "Manta pesada",
            Description = "Dormir biennn"
        })
    elseif index == 7  then
        return Card:New({       
            Title = "Comfort",
            Description = "+4 Comodidad"
        })
    elseif index == 8  then
        return Card:New({       
            Title = "Café",
            Description = "+2 Energía"
        })
    end

    local msg = ""

    if index == nil then
        msg = "Index is nil" .. " " .. self
    end

    return Card:New({
        Cost = 1,
        Title = "Esto no existe",
        Description = msg,
        Tags = {"Error"},
        OnPlay = function (state)
            self.Cost = state.PlayerCount * 2
        end

    })
end


return Character