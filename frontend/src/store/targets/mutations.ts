import { MutationTree } from 'vuex';
import { TargetsState, Species }  from './types'

export const mutations: MutationTree<TargetsState> = {
    setCandidates(state, payload: Species[]) {
        state.candidates = payload;
    }
}
