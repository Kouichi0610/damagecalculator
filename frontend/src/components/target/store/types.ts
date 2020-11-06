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

export interface TargetState {
  name: string;
  types: string[];
  species: Species;
  weight: number;
  abilities: string[];
  currentAbility: string;

  nature: string;
  individuals: Individuals;
  basePoints: BasePoints;

  // 基礎ポイント0～252(4刻み)
  hppattern: number[];
  atpattern: number[];
  dfpattern: number[];
  sapattern: number[];
  sdpattern: number[];
  sppattern: number[];

  //moves: string[];
}

export class StatsPatternArgs {
  level: number;
  name: string;
  nature: string;
  individual: string;
  constructor(level: number, name: string, nature: string, hp: number, at: number, df: number, sa: number, sd: number, sp: number) {
    this.level = level;
    this.name = name;
    this.nature = nature;
    this.individual = "Max";
    if (at == 0) {
      this.individual = "Weakest";
    }
    if (sp == 0) {
      this.individual = "Slowest";
    }
  }
}

export interface Species {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
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
