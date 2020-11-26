import { MutationTree } from 'vuex';
import { AttackerState, Move } from './types';
import { Result } from './sendDamage'

export const mutations: MutationTree<AttackerState> = {
  reset(state) {
    state.moveA = '';
    state.moveB = '';
    state.moveC = '';
    state.moveD = '';
    state.result = [];
  },
  setCurrent(state, current: number) {
    state.currentIndex = current;
  },
  setMove(state, move: string) {
    switch (state.currentIndex) {
      case 0:
        state.moveA = move;
        break;
      case 1:
        state.moveB = move;
        break;
      case 2:
        state.moveC = move;
        break;
      case 3:
        state.moveD = move;
        break;
    }
  },
  setAttacker(state, payload: string) {
    state.attacker = payload;
  },
  setMoves(state, payload: Move[]) {
    state.physicals = [];
    state.specials = [];
    for (var i = 0; i < payload.length; i++) {
      let move = payload[i];
      if (move.isPhysical()) {
        state.physicals.push(move);
      }
      if (move.isSpecial()) {
        state.specials.push(move);
      }
    }
  },
  setResults(state, payload: Result[]) {
    state.result = payload;
  }
}
