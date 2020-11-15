import { StatePatterns } from './statePattern'
import { Species } from './species'
/*
  対象ポケモンに関するStore
  TODO:store主体だと手続き型になるので機能クラスにアクセスする形式に変更
  TODO:target/store改修
  TODO:action.vueからアクセス
  TODO:moveResult.vueに現在の技(+index)を保管
  TODO:defenderDamagesをこっち側に
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

export interface Individuals {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
export interface BasePoints {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
