/*
    能力値計算
*/
import { level } from './level'
import { individuals } from './individuals'
import { basePoints } from './basepoints';
import { species } from './species';
import { nature } from './nature';

/*
func calcHP(l Level, s Species, i Individual, b BasePoint) uint {
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := float32(x)*float32(l)/100.0 + float32(l) + 10
	return uint(y)
}
func calcFlat(l Level, s Species, i Individual, b BasePoint) uint {
	x := uint(s)*2 + uint(i) + uint(b)/4
	y := uint(float32(x)*float32(l))/100 + 5
	return y
}
*/
function calcHP(l, s, i, b) {
    let x = s * 2 + i + b/4;
    let y = x * l / 100.0 + l + 10;
    return Math.floor(y);
}
function calcStats(l, s, i, b, n) {
  let x = s * 2 + i + b/4;
  let y = x * l / 100.0 + 5;
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
      return state.level.level;
    },
    attack: (state) => {
      return state.getters.nature.attack;
    },
    defense: () => {
      return 150;
      //return state.getters.nature.defense;
    },
    spAttack: (state, getters) => {
      return getters.defense + 100;
      //return getters['nature/spAttack'];
    },
    spDefense: (state) => {
      return state.getters.nature.spDefense
    },
    speed: (state) => {
      return state.getters.nature.speed;
    },
    toString (state, getters) {
      let statsStr = 'Stats:' + getters.spAttack;

      return 'Level:' + state.level.level
       + getters['species/toString']
       + getters['individuals/toString']
       + getters['basePoints/toString']
       + getters['nature/currentName']
       + statsStr;
    },
  },
}

