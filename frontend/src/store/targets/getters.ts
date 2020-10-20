import { GetterTree } from 'vuex';
import { TargetsState } from './types';
import { RootState } from '@/store/types';

export const getters: GetterTree<TargetsState, RootState> = {
    targetSpecies: (state: TargetsState) => {
        return [
            state.target.hp,
            state.target.attack,
            state.target.defense,
            state.target.spAttack,
            state.target.spDefense,
            state.target.speed,
        ];
    },
    targetIndividuals : (state: TargetsState) => {
        return [
            state.individuals.hp,
            state.individuals.attack,
            state.individuals.defense,
            state.individuals.spAttack,
            state.individuals.spDefense,
            state.individuals.speed,
        ];
    },
    targetBasePoints: (state: TargetsState) => {
        return [
            state.basePoints.hp,
            state.basePoints.attack,
            state.basePoints.defense,
            state.basePoints.spAttack,
            state.basePoints.spDefense,
            state.basePoints.speed,
        ];
    },
    candidatesCount: (state: TargetsState) => {
        return state.candidates.length;
    },
    candidates: (state: TargetsState) => {
        return state.candidates;
    },
    defaultTypes: (state: TargetsState) => {
        return state.defaultTypes;
    },
    defaultTotal: (state: TargetsState) => {
        return state.defaultTotal;
    }
}

