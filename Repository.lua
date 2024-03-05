local Card = require("Card")

local Repository = {
    DefaultActivityDeckIndexes = {0,0,1,1,0,0,1,1,0,0,1,1,0,0,1,1,0,0,1,1}
}

function Repository:GetActivityDeck()
    return self.DefaultActivityDeckIndexes
end

function Repository:GetActivity (index)
    if index == 0 then
        return Card:New({
            Cost = 2,
            Title = "Desorden!",
            Description = "Dejaron la cagá en la casa",
            Tags = {"Desorden", "Rutina"},
            OnPlay = function (state)
                self.Cost = state.PlayerCount * 2
            end

        })
    elseif index == 1  then
        return Card:New({
            Cost = 3,
            Title = "Sacar la basura",
            Description = "El camión pasa hoy!",
            Tags = {"Rutina", "Olfativo"},
            OnPlay = function (state)
                self.Cost = state.PlayerCount * 2
            end

        })
    elseif index == 2  then
        return Card:New({
            Cost = 2,
            Title = "Pagar las cuentas",
            Description = "Capitalism sucks",
            Tags = {"Rutina"},
            OnPlay = function (state)
                self.Cost = state.PlayerCount * 2
            end

        })
    elseif index == 3  then
        return Card:New({
            Cost = 2,
            Title = "Gotera",
            Description = "Al ritmo de la procastinación",
            Tags = {"Inesperado", "Técnico"},
            OnPlay = function (state)
                self.Cost = state.PlayerCount * 2
            end

        })
    elseif index == 4  then
        return Card:New({
            Cost = 4,
            Title = "Wifi Down!",
            Description = "Todos obtienen -2 comodidad",
            Tags = {"Inesperado", "Técnico"},
            OnPlay = function (state)
                self.Cost = state.PlayerCount * 2
            end

        })
    end
    return Card:New({
        Cost = 0,
        Title = "Esto no existe",
        Description = "Fuera del indice de cartas de actividad",
        Tags = {"Error"},
        OnPlay = function (state)
            self.Cost = state.PlayerCount * 2
        end

    })
end

return Repository