import * as nature from "../../domain/nature"

// リストフィルター
// この条件を満たしているものだけ通す
export interface Filter {
    types: string;  // タイプ
    total: number;  // 種族値合計
}

export interface Species {
    name: string;
    types: string;
    hp: number;
    attack: number;
    defense: number;
    spAttack: number;
    spDefense: number;
    speed: number;
}
export interface Individuals {
    hp: number;
    attack: number;
    defense: number;
    spAttack: number;
    spDefense: number;
    speed: number;
}
export interface BasePoints {
    hp: number;
    attack: number;
    defense: number;
    spAttack: number;
    spDefense: number;
    speed: number;
}

export interface TargetsState {
    target: Species;
    nature: nature.INature;
    individuals: Individuals;
    basePoints: BasePoints;

    defaultTotal: number,
    defaultTypes: string,
    candidates: Species[];
}
