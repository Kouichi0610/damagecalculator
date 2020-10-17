import { MutationTree } from 'vuex';
import { TargetsState, Species }  from './types'

export const mutations: MutationTree<TargetsState> = {
    setCandidates(state, payload: Species[]) {
        state.candidates = payload;

        for (var i = 0; i < state.candidates.length; i++) {
            let s = state.candidates[i];
            var str = '';

            str += s.name;
            str += ' Types:' + s.types;
            str += ' HP:' + s.hp;
            str += ' AT:' + s.attack;
            str += ' DF:' + s.defense;
            str += ' SA:' + s.spAttack;
            str += ' SD:' + s.spDefense;
            str += ' SP:' + s.speed;

            console.log(str)
        }
    }
}
