# TechDice - a Dice roller for Technoir

## Technoir dice

[Technoir](https://www.technoirrpg.com/) is an RPG that uses six-sided dice for its resolution mechanic. The dice pool consists of several types of dice:
- `Action dice`: The player rolling gets a number of `Action dice` equal to their rating in the applicable Verb.
- `Push dice`: The player starts with three available `Push dice`. The player can gain or lose them during play. The player can choose to add `Push dice` to the pool based on using an applicable positive adjective, object, or tag.
- `Hurt dice`: The player must add hurt dice to the pool for every negative adjective they have.

The dice pool is then rolled. Any `Hurt dice` cancel out `Action dice` or `Push dice` with the same number showing. The highest number remaining is the result. If there are multiple dice remaining showing that number, the result is .1 higher (to break ties).

## Dice rolling syntax

The dice rolling syntax is:
```
!tech [Verb] push [Push] hurt [Hurt]
```

The `push` and `hurt` modifiers can be dropped if they are 0.

## Future work

In the future, the bot could also track push dice inventory for players.
