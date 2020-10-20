// チェック対象 候補一覧
import { Module } from 'vuex';
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetsState } from './types';
import { RootState } from '@/store/types';
import * as nature from "../../domain/nature"

// TODO:他stateとのやり取り
export const state: TargetsState = {
    // 現在選択しているポケモン
    target: {name: '', types: '', hp: 0, attack: 0, defense: 0, spAttack: 0, spDefense: 0, speed: 0},
    nature: nature.NatureFactory('てれや'),
    individuals: {hp: 0, attack: 0, defense: 0, spAttack: 0, spDefense: 0, speed: 0},
    basePoints: {hp: 0, attack: 0, defense: 0, spAttack: 0, spDefense: 0, speed: 0},

    // 候補選択
    defaultTotal: 500,
    defaultTypes: 'ノーマル',
    candidates: [],

}

const namespaced: boolean = true;

export const targets: Module<TargetsState, RootState> = {
    namespaced,
    state,
    getters,
    actions,
    mutations,
}
