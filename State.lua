local repo = require("Repository")

local State = {
	Check = "Check",
	Players = {},
	board = {
		drawDeck = {},
		discardDeck = {},
		activeCards = {},
	},
}

function State:InitDefault()
	self.board.drawDeck = repo:GetActivityDeck()
end
function State:AddPlayer(player)
	table.insert(self.Players, player)
end
function State:ActivityPileSize()
	return #self.board.drawDeck
end
function State:DrawActivityCards(n)
	if type(n) == "number" then
		for i = 1, n do
			table.insert(self.board.activeCards, repo:GetActivity(self.board.drawDeck[1]))
			table.remove(self.board.drawDeck, 1)
		end
	end
end
function State:StartPlayerTurn()
	self.Players[1].Deck:DrawCard()
end
function State:PlayPlayerCard(playerIndex, cardIndex)
	-- check if player exists
	local player = self.Players[playerIndex]
	if player == nil then
		return false, "Player index doesn´t exist"
	end

	-- check if player turn is not ended
	if player.turnEnded == false then
		return false, "Player turn ended"
	end

	-- check if card index is valid
	local card = player.deck.hand[cardIndex]
	if card == nil then
		return false, "Card index is not valid"
	end

	-- check if card is valid to play
	if player:CanPlay(cardIndex) == false then
		return false, "Can't play that card"
	end

	-- play card, send update to all
	State:SendPlayPlayerCard(playerIndex, cardIndex)

	return true, "Valid Card"
end

return State

