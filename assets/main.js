var app = new Vue({
  el: '#app',
  data: {
    saveData: {
      date: new Date(),
      caches: [
        {
          "initialValue": 50,
          "unitCost": 1,
          "name": "1yen",
          "value": 0
        }, {
          "initialValue": 50,
          "unitCost": 5,
          "name": "5yen",
          "value": 0
        }, {
          "initialValue": 50,
          "unitCost": 10,
          "name": "10yen",
          "value": 0
        }
      ],
      sales: 0,
      otherServices: [{
        "unitCost": 500,
        "n": 0,
        "isPositive": true,
        "name": "Rabies"
      }],
      unpaids: [0],
      ins: [0],
      outs: [0],
      others: [0],
    }
  },
  methods: {
    appendUnpaid: function() {
      this.saveData.unpaids.push(0)
    },
    appendIns: function() {
      this.saveData.ins.push(0)
    },
    appendOuts: function() {
      this.saveData.outs.push(0)
    },
    appendOthers: function() {
      this.saveData.others.push(0)
    }
  },
  computed: {
    diff: function() {
      let sum = 0
      for (let cache of this.saveData.caches) {
        sum += cache.unitCost * (cache.value - cache.initialValue)
      }
      sum -= this.saveData.sales
      for (let otherService of this.saveData.otherServices) {
        sum += otherService.unitCost * otherService.n
      }
      for (let unpaid of this.saveData.unpaids) {
        sum += unpaid
      }
      for (let inMoney of this.saveData.ins) {
        sum += inMoney
      }
      for (let outMoney of this.saveData.outs) {
        sum -= outMoney
      }
      for (let other of this.saveData.others) {
        sum += other
      }
      return sum
    }
  }
})

