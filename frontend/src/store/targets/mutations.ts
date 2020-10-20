import { MutationTree } from 'vuex';
import { TargetsState, Species, Filter }  from './types'

export const mutations: MutationTree<TargetsState> = {
    setTarget(state, payload: Species) {
        state.target = payload;
    },
    setCandidates(state, payload: Species[]) {
        state.candidates = payload;
    },
    setDefaultTypes(state, payload: string) {
        state.defaultTypes = payload;
    },
    setDefaultTotal(state, payload: number) {
        state.defaultTotal = payload;
    }
}
