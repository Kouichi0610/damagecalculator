/*
    能力値計算
*/
import { level } from './level'
import { individuals } from './individuals'
import { basePoints } from './basepoints';
import { species } from './species';
import { nature } from './nature';

function calcHP(l, s, i, b) {
    let x = s * 2 + i + b/4;
    let y = x * l / 100.0 + l + 10;
    return Math.floor(y);
}
function calcStats(l, s, i, b, n) {
  let x = s * 2 + i + b/4;
  let y = Math.floor(x * l / 100.0 + 5);
  let z = y * n;
  return Math.floor(z);
}

export const stats = {
  namespaced: true,
  modules: {
    level: level,
    species: species,
    individuals: individuals,
    basePoints: basePoints,
    nature: nature,
  },
  state: {
  },
  mutations: {
  },
  getters: {
    hp: (state) => {
      return calcHP(state.level.level, state.species.hp, state.individuals.hp, state.basePoints.hp);
    },
    attack: (state, getters) => {
      return calcStats(state.level.level, state.species.at, state.individuals.at, state.basePoints.at, getters['nature/attack']);
    },
    defense: (state, getters) => {
      return calcStats(state.level.level, state.species.df, state.individuals.df, state.basePoints.df, getters['nature/defense']);
    },
    spAttack: (state, getters) => {
      return calcStats(state.level.level, state.species.sa, state.individuals.sa, state.basePoints.sa, getters['nature/spAttack']);
    },
    spDefense: (state, getters) => {
      return calcStats(state.level.level, state.species.sd, state.individuals.sd, state.basePoints.sd, getters['nature/spDefense']);
    },
    speed: (state, getters) => {
      return calcStats(state.level.level, state.species.sp, state.individuals.sp, state.basePoints.sp, getters['nature/speed']);
    },
    toString (state, getters) {
      let statsStr = 'Stats HP:' + getters.hp
      + ' AT:' + getters.attack
      + ' DF:' + getters.defense
      + ' SA:' + getters.spAttack
      + ' SD:' + getters.spDefense
      + ' SP:' + getters.speed

      return 'Level:' + state.level.level
       + getters['species/toString']
       + getters['individuals/toString']
       + getters['basePoints/toString']
       + getters['nature/currentName']
       + statsStr;
    },
  },
}

