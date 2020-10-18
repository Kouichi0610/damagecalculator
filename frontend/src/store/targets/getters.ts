import { GetterTree } from 'vuex';
import { TargetsState } from './types';
import { RootState } from '@/store/types';

export const getters: GetterTree<TargetsState, RootState> = {
    candidatesCount: (state: TargetsState) => {
        return state.candidates.length;
    },
    candidates: (state: TargetsState) => {
        return state.candidates;
    },
}

