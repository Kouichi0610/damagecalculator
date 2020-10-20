import { MutationTree } from 'vuex';
import { TargetsState, Species, Filter, Individuals, BasePoints }  from './types'

export const mutations: MutationTree<TargetsState> = {
    setTarget(state, payload: Species) {
        state.target = payload;
    },
    setTargetIndividuals(state, payload: Individuals) {
        state.individuals = payload;
    },
    setTargetBasePoints(state, payload: BasePoints) {
        state.basePoints = payload;
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
