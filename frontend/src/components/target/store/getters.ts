import { GetterTree } from 'vuex';
import { TargetState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<TargetState, RootState> = {
  hasTarget: (state: TargetState) => {
    return state.name != '';
  },
  level: (state: TargetState, getters, rootState: RootState) => {
    return rootState.level;
  },
  name: (state: TargetState) => {
    return state.name;
  },
  types: (state: TargetState) => {
    return state.types;
  },
  species: (state: TargetState) => {
    return state.species;
  },
  weight: (state: TargetState) => {
    return state.weight;
  },
  nature: (state: TargetState) => {
    return state.nature;
  },
  individuals: (state: TargetState) => {
    return state.individuals;
  },
  basePoints: (state: TargetState) => {
    return state.basePoints;
  },
  hp (state: TargetState) {
    if (state.hppattern.length == 0) return 0;
    let idx = state.basePoints.hp/4;
    return state.hppattern[idx];
  },
  attack (state: TargetState) {
    if (state.atpattern.length == 0) return 0;
    let idx = state.basePoints.at/4;
    return state.atpattern[idx];
  },
  defense (state: TargetState) {
    if (state.dfpattern.length == 0) return 0;
    let idx = state.basePoints.df/4;
    return state.dfpattern[idx];
  },
  spAttack (state: TargetState) {
    if (state.sapattern.length == 0) return 0;
    let idx = state.basePoints.sa/4;
    return state.sapattern[idx];
  },
  spDefense (state: TargetState) {
    if (state.sdpattern.length == 0) return 0;
    let idx = state.basePoints.sd/4;
    return state.sdpattern[idx];
  },
  speed (state: TargetState) {
    if (state.sppattern.length == 0) return 0;
    let idx = state.basePoints.sp/4;
    return state.sppattern[idx];
  },
}
