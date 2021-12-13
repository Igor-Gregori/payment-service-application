import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { join } from 'path/posix';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { Order } from './orders/entities/order.entity';
import { OrdersModule } from './orders/orders.module';
import { AccountsModule } from './accounts/accounts.module';
import { Account } from './accounts/entities/account.entity';

@Module({
  imports: [
    OrdersModule,
    SequelizeModule.forRoot({
      dialect: 'sqlite',
      host: join(__dirname, 'database.sqlite'),
      autoLoadModels: true,
      models: [Order, Account],
      sync: {
        alter: true,
        // force: true        
      }
    }),
    AccountsModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule { }
