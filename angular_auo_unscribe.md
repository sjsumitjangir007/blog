Auto unsubscribe in Angular 🔥
Ania Sniadek ⚡️
Dev Genius
Ania Sniadek ⚡️

·
Follow

Published in
Dev Genius

·
3 min read
·
Feb 2
504


8






Photo by Joshua Aragon on Unsplash
As we know subscriptions that have been created in a component won’t be destroyed along with the component. When we are using observables in RxJS and want to subscribe to them, we must also remember to unsubscribe to prevent creating a memory leak or to abort HTTP requests in order to avoid unwanted calls. I have a solution for you so that you won’t have to remember about it anymore 🙌.

But first the basics 🙂
There are several options that we can use to unsubscribe our subscriptions.

1️⃣ The first of them is just simply create an array and push your subscriptions there and in the ngOnDestroy() lifecycle method just unsubscribe them all. Here is an example:

@Component(...)
export class AppComponent implements OnInit, OnDestroy {
  subscriptions: Subscription[] = [];

  ngOnInit(): void {
    this.subscriptions.push(of('example value').subscribe());
  }

  ngOnDestroy(): void {
    this.subscriptions.forEach((subscription: Subscription) =>
      subscription.unsubscribe()
    );
  }
}

2️⃣ The second very similar way is to create a variable of type Subscription and add subscriptions to it. In my opinion, this solution is better than the one above because you only need to use the unsubscribe() method and all added subscriptions will be destroyed, you don’t need to use the loop and unsubscribe for each one of them. It will look like this:
@Component(...)
export class AppComponent implements OnInit, OnDestroy {
  subscription: Subscription = new Subscription();

  ngOnInit(): void {
    this.subscription.add(of('example value').subscribe());
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }
}

3️⃣ The last example and my favorite one ⭐️ is to use variable of type Subject<void>. There is no need to push or add any subscriptions in here. Instead of this we are using takeUntil() operator from the RxJS library, but also we need to remember to destroy our subscriptions in ngOnDestroy(). For this solution see code below 👇
@Component(...)
export class AppComponent implements OnInit, OnDestroy {
  destroy$: Subject<void> = new Subject<void>();

  ngOnInit(): void {
    of('example value').pipe(takeUntil(this.destroy$)).subscribe();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}

Great! 🤩

We have a couple of solutions to keep an eye on our subscriptions so there are no memory leaks, but what if we can do better?

What if we don’t have to remember to unsubscribe and always write the same code in ngOnDestroy(), because it will be done automatically.

You don’t believe it? 👀 See an example of this awesome property decorator below 👇

@AutoDestroy Decorator 🔥
This solution only works with the third method above. If you prefer to use one of the first two methods, that’s fine, just rewrite this decorator to execute the methods used in ngOnDestroy().

The full code of @AutoDestroy decorator will look like this:
export function AutoDestroy(component: any, key: string | symbol): void {
  const originalOnDestroy = component.ngOnDestroy;

  component.ngOnDestroy = function () {
    if (originalOnDestroy) {
      originalOnDestroy.call(this);
    }
    this[key].next();
    this[key].complete();
  }
}

If you want to use this decorator in component, you can do this like so:

@Component(...)
export class AppComponent implements OnInit {
  @AutoDestroy destroy$: Subject<void> = new Subject<void>();

  ngOnInit(): void {
    of('example value').pipe(takeUntil(this.destroy$)).subscribe();
  }
}

So how does it work? It is very simple!

Check it out below 👇

How it works? 🤔
To use @AutoDestroy decorator we just need to add it before our destroy$ variable of type Subject<void>. Thanks to this, the component and the name of this variable will be passed to our decorator. The main purpose of this decorator is to overwrite the ngOnDestroy() lifecycle method with its own. It is very important to check if this component doesn’t already have its own ngOnDestroy(), because without it we will lose the original call. So if the component has its own ngOnDestroy() we just simply call it here, then destroy our subscriptions, and that’s all.

Easy, right? 😎

You don’t have to remember to unsubscribe anymore, just use @AutoDestroy property decorator 🔥.

Wrapping up
I hope this article will help you with subscriptions problems. If you want to test it check it out on StackBlitz ⚡️ or look through the source code on Github 🎓.

If you’ve enjoyed reading this, please give a few claps 👏 . Feel free to follow me and to leave a comment with your feedback 👇.


![angular_auto_unsubscribe_Web capture_20-11-2023_182456_blog devgenius io](https://github.com/sjsumitjangir007/sumitj_blog/assets/17833671/404f4bba-f8f5-4996-b700-280f359b6222)
