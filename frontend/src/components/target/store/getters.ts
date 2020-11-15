import { GetterTree } from 'vuex';
import { TargetState } from './types';
import { RootState } from '@/store/types'
import { Species } from './species'

export const getters: GetterTree<TargetState, RootState> = {
  level: (state: TargetState, getters, rootState: RootState): number => {
    return rootState.level;
  },
  species: (state: TargetState): Species => {
    return state.species;
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
    return state.statePatterns.hp(state.basePoints.hp);
  },
  attack (state: TargetState) {
    return state.statePatterns.attack(state.basePoints.at);
  },
  defense (state: TargetState) {
    return state.statePatterns.defense(state.basePoints.df);
  },
  spAttack (state: TargetState) {
    return state.statePatterns.spAttack(state.basePoints.sa);
  },
  spDefense (state: TargetState) {
    return state.statePatterns.spDefense(state.basePoints.sd);
  },
  speed (state: TargetState) {
    return state.statePatterns.speed(state.basePoints.sp);
  },
  currentAbility (state: TargetState) {
    return state.currentAbility;
  }
}
