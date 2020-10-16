// 個体値
export const individuals = {
  namespaced: true,
  state: {
      hp: 31,
      at: 31,
      df: 31,
      sa: 31,
      sd: 31,
      sp: 31,
  },
  getters: {
    toString (state) {
      return '個体値:' + state.hp + ' ' + state.at + ' ' + state.df + ' ' + state.sa + ' ' + state.sd + ' ' + state.sp;
    }
  },
  mutations: {
    toMax (state) {
      state.at = 31;
      state.sp = 31;
    },
    toWeakest (state, isWeakest) {
      state.at = isWeakest ? 0 : 31;
    },
    toSlowest (state, isSlowest) {
      state.sp = isSlowest ? 0: 31;
    }
  }
}
