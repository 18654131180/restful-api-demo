package impl

import (
	"context"
	"fmt"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/sqlbuilder"
)

// 业务处理层(Controller层)
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (
	*host.Host, error) {
	// 直接打印日志
	i.l.Named("Creaet").Debug("create host")
	i.l.Info("create host")
	// // 带Format的日志打印, fmt.Sprintf()
	i.l.Debugf("create host %s", ins.Name)
	// // 携带额外meta数据, 常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 默认值填充
	ins.InjectDefault()

	// 有dao模块 负责 把对象入库
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	b := sqlbuilder.NewBuilder(QueryHostSQL)

	if req.Keywords != "" {
		// (r.`name`='%' OR r.description='%' OR r.private_ip='%' OR r.public_ip='%')
		//  10.10.1, 接口测试
		b.Where("r.`name`LIKE ? OR r.description LIKE ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	b.Limit(req.OffSet(), req.GetPageSize())
	querySQL, args := b.Build()
	i.l.Debugf("query sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := host.NewHostSet()
	for rows.Next() {
		// 没扫描一行,就需要读取出来
		// h.cpu, h.memory, h.gpu_spec, h.gpu_amount, h.os_type, h.os_name, h.serial_number
		ins := host.NewHost()
		if err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
			&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
			&ins.Account, &ins.PublicIP, &ins.PrivateIP,
			&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
		); err != nil {
			return nil, err
		}
		set.Add(ins)
	}

	// total统计
	countSQL, args := b.BuildCount()
	i.l.Debugf("count sql: %s, args: %v", countSQL, args)
	countStmt, err := i.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer countStmt.Close()
	if err := countStmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}

	return set, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (
	*host.Host, error) {
	b := sqlbuilder.NewBuilder(QueryHostSQL)
	b.Where("r.id = ?", req.Id)

	querySQL, args := b.Build()
	i.l.Debugf("describe sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ins := host.NewHost()
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
		&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
		&ins.Account, &ins.PublicIP, &ins.PrivateIP,
		&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (
	*host.Host, error) {
	// 获取已有对象
	ins, err := i.DescribeHost(ctx, host.NewDescribeHostRequestWithId(req.Id))
	if err != nil {
		return nil, err
	}

	i.l.Debugf("host in db: %s", *ins)
	// 根据更新的模式, 更新对象
	switch req.UpdateMode {
	case host.UPDATE_MODE_PUT:
		if err := ins.Put(req.Host); err != nil {
			return nil, err
		}
		// 整个对象的局部更新
	case host.UPDATE_MODE_PATCH:
		if err := ins.Patch(req.Host); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("update_mode only requred put/patch")
	}

	i.l.Debugf("host update: %s", *ins)
	// 检查更新后的数据是否合法
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 更新数据库里面的数据
	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	// 返回更新后的对象
	return ins, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	return nil, nil
}
