import { ActionTree } from 'vuex';
import { RootState } from '@/store/types';
import { AttackerState, MoveLoader } from './types'
import { state } from './index'

export const actions: ActionTree<AttackerState, RootState> = {
  loadMoves: ({ commit }, attacker: string) => {
    // TODO:stateをexportはまずい気もするが

    // 対象変更時に読み直す
    if (state.attacker == attacker) return;
    commit('reset');
    new MoveLoader().load(attacker)
    .then((moves) => {
      commit('setMoves', moves);
      commit('setAttacker', attacker);
    });
  },
}
