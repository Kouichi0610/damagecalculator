// 基礎ポイント
export const basePoints = {
  namespaced: true,
  state: {
    hp: 0,
    at: 0,
    df: 0,
    sa: 0,
    sd: 0,
    sp: 0,
  },
  getters: {
    toString (state) {
      return '基礎ポイント:' + state.hp + ' ' + state.at + ' ' + state.df + ' ' + state.sa + ' ' + state.sd + ' ' + state.sp;
    }
  },
  mutations: {
    setBasePoints (state, hp, at, df, sa, sd, sp) {
      state.hp = hp;
      state.at = at;
      state.df = df;
      state.sa = sa;
      state.sd = sd;
      state.sp = sp;
    }
  }
}
