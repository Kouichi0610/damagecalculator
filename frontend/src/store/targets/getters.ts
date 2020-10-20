import { GetterTree } from 'vuex';
import { TargetsState } from './types';
import { RootState } from '@/store/types';

export const getters: GetterTree<TargetsState, RootState> = {
    targetSpecies: (state: TargetsState) => {
        return state.target;
    },
    targetStr: (state: TargetsState) => {
        if (state.target == null) {
            return "(none)";
        }
        let t = state.target;
        return '' + t.name + ' ' + t.types + ' HP:' + t.hp + ' AT:' + t.attack + ' DF:' + t.defense + ' SA:' + t.spAttack + ' SD:' + t.spDefense + ' SP:' + t.speed;
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

