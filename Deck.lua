local Deck = {
    hand = {},
    deck = {},
    discards = {},
    active = {},
    repositoryMethod = function(self, index) end,
}

function Deck:New(o)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    return o
end

function Deck:DrawCard()
    print("DrawCard")
    local c = self:repositoryMethod(self.deck[1])
    table.insert(self.hand, c)
    table.remove(self.deck, 1)
    return c
end

function Deck:DiscardCard(index)
    table.insert(self.discardDeck, self.hand[index])
    table.remove(self.hand, index)
end

function Deck:DrawAndPlay(repositoryMethod)
    table.insert(self.board.activeCards, self:repositoryMethod(self.board.drawDeck[1]))
    table.remove(self.board.drawDeck, 1)
end

return Deck