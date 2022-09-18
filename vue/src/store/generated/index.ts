// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Zireael26CheckersCheckers from './zireael26.checkers.checkers'


export default { 
  Zireael26CheckersCheckers: load(Zireael26CheckersCheckers, 'zireael26.checkers.checkers'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}