import {VueRouter} from 'vue-router'
import root from './engines'

console.log(root)
const router = new VueRouter()
router.map(root.routes)

export default router
