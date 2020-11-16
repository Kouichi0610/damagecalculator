import { GetterTree } from 'vuex';
import { TargetState } from './types';
import { RootState } from '@/store/types'
import { Species } from './species'
import { DefenderDamages } from './defenderDamages'

export const getters: GetterTree<TargetState, RootState> = {
  defenderDamages: (state: TargetState, getters, rootState: RootState): DefenderDamages => {
    let attacker = state.species.name;
    let level = rootState.level;
    let individuals = state.individuals.type();
    let basePoints = state.basePoints;
    let ability = state.currentAbility;
    let nature = state.nature;
    let move = 'はかいこうせん';  // TODO:
    let item = '';  // TODO:
    let condition = ''; // TODO:
    let weather = ''; // TODO:
    let field = ''; // TODO:
    return new DefenderDamages(attacker, level, individuals, basePoints, ability, nature, move, item, condition, weather, field);
  },
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
    return state.statePatterns.hp(state.basePoints);
  },
  attack (state: TargetState) {
    return state.statePatterns.attack(state.basePoints);
  },
  defense (state: TargetState) {
    return state.statePatterns.defense(state.basePoints);
  },
  spAttack (state: TargetState) {
    return state.statePatterns.spAttack(state.basePoints);
  },
  spDefense (state: TargetState) {
    return state.statePatterns.spDefense(state.basePoints);
  },
  speed (state: TargetState) {
    return state.statePatterns.speed(state.basePoints);
  },
  currentAbility (state: TargetState) {
    return state.currentAbility;
  }
}
