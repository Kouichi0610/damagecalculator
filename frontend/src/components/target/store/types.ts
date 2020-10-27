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

interface Species {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
interface Individuals {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
interface BasePoints {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
