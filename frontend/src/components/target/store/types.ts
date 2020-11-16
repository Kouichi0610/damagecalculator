import { StatePatterns } from './statePattern'
import { Species } from './species'
import { Individuals } from './individuals'
/*
  対象ポケモンに関するStore
  TODO:action.vueからアクセス
  TODO:moveResult.vueに現在の技(+index)を保管
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

export interface BasePoints {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
