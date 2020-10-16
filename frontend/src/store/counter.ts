// モジュール切り分け試作
// 削除予定
// (ただのオブジェクトなので注意 そのままでstore.gettersのようには使えない)
export const counter = {
    namespaced: true,
    // 状態 他のコンポーネントから直接書き換えることはできない
    state: {
        count: 0,
    },
    // getters 状態の取得(stateの直接取得は可能)
    // 値をそのまま取得するだけならstate.countで良い
    // (computedの中でgettersを呼ばない)
    getters: {
        count: state => state.count,
        isPositive: state => {
            return state.count > 0;
        },
        isNegative: state => {
            return state.count < 0;
        },
        /*
            第一：ローカルモジュール
            第二：getters
            第三：ルートStore(まず使わないが)
        */
        sample (/*state, getters, route*/) {
            // こんな感じで
            // return getters['child/childMethod'];
            // state.gettersのようには使えない
        }
    },
    // 状態の変更
    /*
        利用者側の記述例
        store.commit('counter/add', 15);
    */
    mutations: {
        increment (state) {
            state.count++;
        },
        decrement (state) {
            state.count--;
        },
        add (state, amount) {
            state.count += amount;
        },
        sub (state, amount) {
            state.count -= amount;
        }
    },
    // 外部APIとの連携、非同期処理
    // ここでは1秒後に加減を行っている
    actions: {
        addAsync( {commit}, payload) {
            return new Promise((resolve/*, reject*/) => {
                setTimeout(() => {
                    commit('add', payload.amount);
                    resolve();
                }, 1000);
            });
        },
        subAsync( {commit}, payload) {
            return new Promise((resolve/*, reject*/) => {
                setTimeout(() => {
                    commit('sub', payload.amount);
                    resolve();
                }, 1000);
            });
        }
    },
}


