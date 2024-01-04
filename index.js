/** 
  * Sikiの起動時に一度だけ実行されるスクリプト 
  * フォルダをpluginsへ配置して有効化
  * 
  * @version 0.0.1
 */

/**
 * プラグインのタイプ
 * startupは起動時にmainScriptの内容を一度だけ実行します
 */
module.exports.type = 'startup'

/**
 * プラグインの情報
 */
module.exports.meta = {
  name: 'siki-plugin-startup-counter',
  description: 'This plugin counts the number of times Siki has been started in the past day.',
  version: '1.0.0',
  needVersion: '0.24.7'
}

const { execSync } = require('child_process')
const { join } = require('path')

/**
 * 
 * @param {Object} settings 
 * @param {{[key: string]: any}} settings.config - config.jsに記述された設定内容
 * @param {{[key: string]: any}} settings.userSettings - user.jsに記述された設定内容
 */
module.exports.mainScript = async (settings) => {
  return execSync(join(__dirname, 'counter.out')).toString()
}
