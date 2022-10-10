// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IERC20} from "./IERC20.sol";

// import "@openzeppelin./";

contract TravelSaver {
    /**
     ***** ***** STRUCTS ***** *****
     */

    struct TravelPlan {
        address owner;
        uint256 ID;
        uint256 operatorPlanID;
        uint256 operatorUserID;
        uint256 contributedAmount;
        uint256 createdAt;
        uint256 claimedAt;
        bool claimed;
    }

    struct PaymentPlan {
        uint256 travelPlanID;
        uint256 ID;
        uint256 totalAmount;
        uint256 amountSent;
        uint256 amountPerInterval;
        uint256 totalIntervals;
        uint256 intervalsProcessed;
        uint256 nextTransferOn;
        uint256 interval;
        address sender;
        bool alive;
    }

    /**
     ***** ***** STATE-VARIABLES ***** *****
     */

    // modifier onlyOwner() {
    //     require(msg.sender == owner);
    //     _;
    // }

    address public immutable operatorWallet;
    IERC20 public immutable token;

    uint256 travelPlanCount;
    uint256 paymentPlanCount;

    mapping(uint256 => TravelPlan) public travelPlans;
    mapping(uint256 => PaymentPlan) paymentPlans;
    mapping(uint256 => mapping(address => uint256)) public contributedAmount;

    constructor(address ERC20_, address operatorWallet_) {
        token = IERC20(ERC20_);
        operatorWallet = operatorWallet_;
    }

    /**
     ***** ***** VIEW-FUNCTIONS ***** *****
     */

    function getTravelPlanDetails(uint256 ID)
        external
        view
        returns (TravelPlan memory)
    {
        return travelPlans[ID];
    }

    function getPaymentPlanDetails(uint256 ID)
        external
        view
        returns (PaymentPlan memory)
    {
        return paymentPlans[ID];
    }

    /**
     ***** ***** EVENTS ***** *****
     */

    event CreatedTravelPlan(
        uint256 indexed ID,
        address indexed owner,
        TravelPlan travelPlan
    );

    event ContributeToTravelPlan(
        uint256 indexed ID,
        address indexed contributor,
        uint256 amount
    );
    event ClaimTravelPlan(uint256 indexed ID);

    event Transfer(address indexed from, address indexed to, uint256 amount);

    event CreatedPaymentPlan(
        uint256 indexed ID,
        address indexed owner,
        PaymentPlan paymentPlan
    );

    event CancelPaymentPlan(
        uint256 indexed ID,
        address indexed owner,
        PaymentPlan paymentPlan
    );

    event StartPaymentPlanInterval(
        uint256 indexed ID,
        uint256 indexed callableOn,
        uint256 indexed amount,
        uint256 intervalNo
    );
    event PaymentPlanIntervalEnded(
        uint256 indexed ID,
        uint256 indexed intervalNo
    );
    event EndPaymentPlan(
        uint256 indexed ID,
        address indexed owner,
        PaymentPlan paymentPlan
    );

    /**
     ***** ***** STATE-CHANGING-EXTERNAL-FUNCTIONS ***** *****
     */

    function createTravelPaymentPlan(
        uint256 operatorPlanID_,
        uint256 operatorUserID_,
        uint256 amountPerInterval,
        uint256 totalIntervals,
        uint256 intervalLength
    ) external returns (uint256 travelPlanID, uint256 paymentPlanID) {
        travelPlanID = createTravelPlan(operatorPlanID_, operatorUserID_);
        paymentPlanID = createPaymentPlan(
            travelPlanID,
            amountPerInterval,
            totalIntervals,
            intervalLength
        );
        return (travelPlanID, paymentPlanID);
    }

    function createTravelPlan(uint256 operatorPlanID_, uint256 operatorUserID_)
        public
        returns (uint256)
    {
        travelPlanCount += 1;

        travelPlans[travelPlanCount] = TravelPlan({
            owner: msg.sender,
            ID: travelPlanCount,
            operatorPlanID: operatorPlanID_,
            operatorUserID: operatorUserID_,
            contributedAmount: 0,
            createdAt: block.timestamp,
            claimedAt: 0,
            claimed: false
        });

        emit CreatedTravelPlan(
            travelPlanCount,
            msg.sender,
            travelPlans[travelPlanCount]
        );
        return travelPlanCount;
    }

    function contributeToTravelPlan(uint256 ID, uint256 amount) external {
        TravelPlan storage plan = travelPlans[ID];
        require(block.timestamp >= plan.createdAt, "doesn't exist");
        require(!plan.claimed, "claimed");

        plan.contributedAmount += amount;

        contributedAmount[ID][msg.sender] += amount;
        token.transferFrom(msg.sender, address(this), amount);

        emit ContributeToTravelPlan(ID, msg.sender, amount);
        emit Transfer(msg.sender, address(this), amount);
    }

    function claimTravelPlan(uint256 ID) external {
        TravelPlan storage plan = travelPlans[ID];
        require(plan.owner == msg.sender, "not owner");
        require(plan.contributedAmount > 0, "nothing saved");
        require(!plan.claimed, "plan claimed");

        token.transfer(operatorWallet, plan.contributedAmount);
        plan.claimed = true;
        plan.claimedAt = block.timestamp;
        emit ClaimTravelPlan(ID);
        emit Transfer(address(this), operatorWallet, plan.contributedAmount);
    }

    function createPaymentPlan(
        uint256 _travelPlanID,
        uint256 amountPerInterval,
        uint256 totalIntervals,
        uint256 intervalLength
    ) public returns (uint256) {
        uint256 totalToTransfer = amountPerInterval * totalIntervals;
        require(
            IERC20(token).allowance(msg.sender, address(this)) >=
                totalToTransfer,
            "IERC20: Insuff Approval"
        );
        uint256 id = ++paymentPlanCount;

        paymentPlans[id] = PaymentPlan({
            travelPlanID: _travelPlanID,
            ID: id,
            totalAmount: totalIntervals * amountPerInterval,
            amountSent: 0,
            amountPerInterval: amountPerInterval,
            totalIntervals: totalIntervals,
            intervalsProcessed: 0,
            nextTransferOn: 0,
            interval: intervalLength,
            sender: msg.sender,
            alive: true
        });
        _startInterval(id);

        emit CreatedPaymentPlan(id, msg.sender, paymentPlans[id]);

        return id;
    }

    function cancelPaymentPlan(uint256 ID) external {
        require(msg.sender == paymentPlans[ID].sender, "only plan owner");
        _endPaymentPlan(ID);

        emit CancelPaymentPlan(ID, msg.sender, paymentPlans[ID]);
    }

    function runInterval(uint256 ID) external {
        _fulfillPaymentPlanInterval(ID);
    }

    function runIntervals(uint256[] memory IDs) external {
        for (uint256 i = 0; i < IDs.length; i++) {
            _fulfillPaymentPlanInterval(IDs[i]);
        }
    }

    /**
     ***** ***** STATE-CHANGING-PRIVATE-FUNCTIONS ***** *****
     */

    function _startInterval(uint256 ID) internal {
        PaymentPlan memory plan = paymentPlans[ID];
        uint256 callableOn = paymentPlans[ID].interval + block.timestamp;
        uint256 intervalNumber = plan.intervalsProcessed + 1;
        paymentPlans[ID].nextTransferOn = callableOn;

        emit StartPaymentPlanInterval(
            ID,
            callableOn,
            plan.amountPerInterval,
            intervalNumber
        );
    }

    function _endPaymentPlan(uint256 ID) internal {
        PaymentPlan memory plan = paymentPlans[ID];
        paymentPlans[ID].alive = false;
        emit EndPaymentPlan(ID, plan.sender, plan);
    }

    function _contributeToTravelPlan(
        uint256 ID,
        uint256 amount,
        address caller
    ) internal {
        TravelPlan storage plan = travelPlans[ID];
        require(block.timestamp >= plan.createdAt, "doesn't exist");
        require(!plan.claimed, "claimed");

        plan.contributedAmount += amount;

        contributedAmount[ID][caller] += amount;
        token.transferFrom(caller, address(this), amount);

        emit ContributeToTravelPlan(ID, caller, amount);
        emit Transfer(caller, address(this), amount);
    }

    function _fulfillPaymentPlanInterval(uint256 ID) internal {
        PaymentPlan memory plan = paymentPlans[ID];

        uint256 amountToTransfer = plan.amountPerInterval;
        address sender = plan.sender;
        uint256 interval = plan.intervalsProcessed + 1;
        require(plan.nextTransferOn <= block.timestamp, "too early");
        require(plan.alive, "plan ended");

        // Check conditions here with an if clause instead of require, so that integrators dont have to keep track of balances
        if (
            token.balanceOf(sender) >= amountToTransfer &&
            token.allowance(sender, address(this)) >= amountToTransfer
        ) {
            _contributeToTravelPlan(
                plan.travelPlanID,
                amountToTransfer,
                sender
            );

            paymentPlans[ID].amountSent += amountToTransfer;
            paymentPlans[ID].intervalsProcessed = interval;

            emit PaymentPlanIntervalEnded(ID, interval);

            if (interval < plan.totalIntervals) {
                _startInterval(ID);
            } else {
                _endPaymentPlan(ID);
            }
        }
    }
}
