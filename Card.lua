
local Card = {
    Cost = 2,
    Title = "Título Genérico",
    Description = "Description",
    Tags = {},
}
    
function Card:OnPlay(state)
    Cost = state.PlayerCount * 2;
end

function Card:OnExitPlay(state)

end
function Card:ToString()
    return self.Cost .. " " .. self.Title .. ": " .. self.Description
end
function Card:New (o)
    o = o or {}   -- create object if user does not provide one
    setmetatable(o, self)
    self.__index = self
    return o
end
return Card