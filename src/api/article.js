import request from '@/utils/request'
// 分类：获取分类
export const artGetChannelsService = () => request.get('/my/cate/list')
// 分类：添加分类
export const artAddChannelService = (data) => request.post('/my/cate/add', data)
// 分类：编辑分类
export const artEditChannelService = (data) =>
  request.put('/my/cate/info', data)
// 分类：删除分类
export const artDelChannelService = (id) =>
  request.delete('/my/cate/del', {
    params: { id }
  })

//表白 ：获取表白列表
export const artGetListService = (params) =>
  request.get('/my/article/list', {
    params
  })

// 表白：添加表白
export const artPublishService = (data) => request.post('/my/article/add', data)

// 表白：获取表白详情
export const artGetDetailService = (id) =>
  request.get('/my/article/info', {
    params: { id }
  })

// 表白：编辑表白接口
export const artEditService = (data) => request.put('/my/article/info', data)

// 表白：删除表白接
export const artDelService = (id) =>
  request.delete('/my/article/info', { params: { id } })
