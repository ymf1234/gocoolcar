// pages/mytrips/mytrips.ts

import { IAppOption } from "../../appoption"

interface Trip {
    shortId: string
    start: string
    end: string
    duration: string
    fee: string
    distance: string
    status: string
}

interface MainItem {
    id: string
    navId: string
    navScrollId: string
    data: Trip
}

interface NavItem {
    id: string
    mainId: string
    label: string
}

interface MainItemQueryResult {
    id: string
    top: number
    dataset: {
        navId: string
        navScrollId: string
    }
}

Page({

    scrollStates: {
        mainItems: [] as MainItemQueryResult[],
    },

    /**
     * 页面的初始数据
     */
    data: {
        indicatorDots: true,
        autoPlay:true,
        interval:3000,
        duration:500,
        circular: true,
        multiItemCount: 1,
        preVMargin: '',
        nextMargin: '',
        vertical: false,
        current: 0,
        promotionItems: [
            {
                img: 'https://img.mukewang.com/5f7301d80001fdee18720764.jpg',
                promotionID: 1,
            },            
            {
                img: 'https://img.mukewang.com/5f6805710001326c18720764.jpg',
                promotionID: 2,
            },
            {
                img: 'https://img.mukewang.com/5f6173b400013d4718720764.jpg',
                promotionID: 3,
            },
            {
                img: 'https://img.mukewang.com/5f7141ad0001b36418720764.jpg',
                promotionID: 4,
            },
        ],
        avatarURL: '',
        tripsHeight: 0,
        mainItems: [] as MainItem[],
        navItems: [] as NavItem[],
        mainScroll: '',
        navCount: 0,
        navSel: '',
        navScroll: '',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    async onLoad() {
        this.populateTrips()
        const userInfo = await getApp<IAppOption>().globalData.userInfo
        this.setData({
            avatarURL: userInfo?.avatarUrl
        })
    },

    populateTrips() {
        const mainItems: MainItem[] = []
        const navItems: NavItem[] = []
        let navSel = ''
        let preNav = ''
        for(let i =0; i<100;i++) {
            const mainId = 'main-' + i
            const navId = 'nav-' + i
            const tripId = (10001 + i).toString()
            if(!preNav) {
                preNav = navId
            }
            mainItems.push({
                id: mainId,
                navId: navId,
                navScrollId: preNav,
                data: {
                    shortId: tripId,
                    start: '上海',
                    end: '河南',
                    distance: '776公里',
                    duration: '9时50分',
                    fee: '800元',
                    status: '已完成',
                } 
            }),

            navItems.push({
                id: navId,
                mainId: mainId,
                label: tripId,
            })
            if (i===0) {
                navSel = navId
            }
            preNav = navId
        }
        this.setData({
            mainItems: mainItems,
            navItems: navItems,
            navSel: navSel,
        }, () => {
            this.prepareScrollStates()
        })
    },

    prepareScrollStates() {
        wx.createSelectorQuery().selectAll(".main-item")
        .fields({
            id: true,
            dataset: true,
            rect: true
        }).exec(res => {
            this.scrollStates.mainItems = res[0]
        })
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {
        wx.createSelectorQuery().select("#heading")
        .boundingClientRect(rect => {
            const height = wx.getSystemInfoSync().windowHeight - rect.height
            this.setData({
                tripsHeight : height,
                navCount: height / 50
            })
        }).exec()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    },

    onSwiperChange(e: any) {
    },

    onPromotionItemTap(e: any) {
        console.log(e)
        const promotionId = e.currentTarget.dataset.promotionId
        if (promotionId) {

        }
    },

    onNavItemTap(e: any) {
        const mainId: string = e.currentTarget?.dataset?.mainId
        const navId: string = e.currentTarget?.id
        if (mainId) {
            this.setData({
                mainScroll : mainId,
                navSel: navId
            })
        }
    },

    onMainScroll(e: any) {
        const top: number = e.currentTarget?.offsetTop + e.detail?.scrollTop
        if (top === undefined) {
            return
        }

        const selitem = this.scrollStates.mainItems.find(
            v => v.top >= top
        )

        if(!selitem) {
            return
        }

        this.setData({
            navSel: selitem.dataset.navId,
            navScroll: selitem.dataset.navScrollId
        })

    },
})