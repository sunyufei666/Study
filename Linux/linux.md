切换到root

```
su -
```

安装

```
# 不会自动补全依赖
sudo dpkg -i package.deb
# 会自动补全依赖
sudo apt install package.deb
# 修复损坏的依赖
sudo apt --fix-broken install
```

卸载

```
# 普通卸载
sudo apt remove package.deb
# 彻底卸载（包括相关配置）
sudo apt purge package.deb
# 卸载已不需要使用的依赖以及相关配置
sudo apt autoremove --purge
# 清理过时的 APT 缓存包
sudo apt autoclean
```

Git

```
// 取消所有代理设置
git config --global --unset http.proxy
git config --global --unset https.proxy
```

