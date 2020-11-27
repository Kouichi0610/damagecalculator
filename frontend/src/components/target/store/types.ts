import { StatePatterns } from './statePattern'
import { Species } from './species'
import { Individuals } from './individuals'
import { BasePoints } from './basePoints'
/*
  対象ポケモンに関するStore
*/
export interface TargetState {
  species: Species;
  currentAbility: string;
  nature: string;
  individuals: Individuals;
  basePoints: BasePoints;

  // 基礎ポイント0～252(4刻み)
  statePatterns: StatePatterns;

  //moves: string[];
}
