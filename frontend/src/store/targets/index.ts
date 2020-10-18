// チェック対象 候補一覧
import { Module } from 'vuex';
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetsState } from './types';
import { RootState } from '@/store/types';

export const state: TargetsState = {
    defaultTotal: 500,
    defaultTypes: 'ノーマル',
    target: {name: '', types: '', hp: 0, attack: 0, defense: 0, spAttack: 0, spDefense: 0, speed: 0},
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
