import { MutationTree } from 'vuex';
import { SpeedState } from './types';
import { SpeedInfo } from './types'

export const mutations: MutationTree<SpeedState> = {
  setTarget(state, payload: string) {
    state.target = payload;
  },
  setLevel(state, payload: number) {
    state.level = payload;
  },
  setList(state, list: SpeedInfo[]) {
    state.otherList = list;
  },
  setTargetSpeed(state, speed: number) {
    state.targetSpeed = speed;
  },
  setTrickRoom(state, payload: boolean) {
    state.trickRoom = payload;
  }
}
