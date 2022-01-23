import { HttpCode, Inject, Injectable } from '@nestjs/common';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectModel } from '@nestjs/sequelize';
import { EmptyResultError } from 'sequelize';
import { AccountStorageService } from 'src/accounts/account-storage/account-storage.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectModel(Order)
    private orderModule: typeof Order,
    private accountStorage: AccountStorageService,
    @Inject('KAFKA_PRODUCER')
    private kafkaProducer: Producer,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    const order = await this.orderModule.create({
      ...createOrderDto,
      account_id: this.accountStorage.account.id,
    });
    this.kafkaProducer.send({
      topic: 'transactions',
      messages: [
        {
          key: 'transactions',
          value: JSON.stringify({
            id: order.id,
            account_id: order.account_id,
            credit_card_number: order.credit_card_number,
            credit_card_name: order.credit_card_name,
            credit_card_expiration_month: (createOrderDto as any)
              .credit_card_expiration_month,
            credit_card_expiration_year: (createOrderDto as any)
              .credit_card_expiration_year,
            credit_card_expiration_cvv: (createOrderDto as any).credit_card_cvv,
            amount: order.amount,
          }),
        },
      ],
    });
    return order;
  }

  findAll() {
    return this.orderModule.findAll({
      where: {
        account_id: this.accountStorage.account.id,
      },
    });
  }

  findOneByAccountId(id: string) {
    return this.orderModule.findOne({
      where: {
        id,
        account_id: this.accountStorage.account.id,
      },
      rejectOnEmpty: new EmptyResultError(`Order with ID ${id} not found`),
    });
  }

  findOne(id: string) {
    return this.orderModule.findByPk(id);
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    // const account = this.accountStorage.account;
    // const order = await (account ? this.findOneByAccountId(id) : this.findOne(id));
    const order = await this.findOne(id);
    return order.update(updateOrderDto);
  }

  @HttpCode(204)
  async remove(id: string) {
    const order = await this.findOneByAccountId(id);
    return order.destroy();
  }
}
