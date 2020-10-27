/*
  対象ポケモンに関するStore

  情報
  種族値、わざ、とくせい、おもさ

  個体、基礎ポイント、性格

  やりたいこと
  すばやさ調整
  攻撃調整
  耐久調整

  TODO:サーバから計算式を渡せないか
*/
import { INature } from "../../../domain/nature";

export interface TargetState {
  name: string;
  types: string[];
  species: Species;
  weight: number;
  abilities: string[];

  nature: INature;
  individuals: Individuals;
  basePoints: BasePoints;
  //moves: string[];
}

export default interface Species {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
export default class Individuals {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
  constructor(hp: number, at: number, df: number, sa: number, sd: number, sp: number) {
    this.hp = hp;
    this.at = at;
    this.df = df;
    this.sa = sa;
    this.sd = sd;
    this.sp = sp;
  }

  changeSlowest(isSlowest: boolean) {
    this.sp = isSlowest ? 0 : 31;
  }
  changeWeakest(isWeakest: boolean) {
    this.at = isWeakest ? 0 : 31;
  }
}
export interface BasePoints {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
