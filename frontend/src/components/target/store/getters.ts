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
  abilities (state: TargetState) {
    return state.abilities;
  },
  currentAbility (state: TargetState) {
    return state.currentAbility;
  }
}
